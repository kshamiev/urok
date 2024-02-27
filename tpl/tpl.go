package tpl

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/radovskyb/watcher"
)

var (
	fileParseExt = map[string]struct{}{".html": {}}
	dirExcluded  = map[string]struct{}{"/temp/": {}, "/assets/": {}}
)

type HTMLGenerator struct {
	functions map[string]interface{}
	tplStore  map[string]*template.Template
	mu        sync.RWMutex
}

func NewTplGenerator(dirTpl string, f map[string]interface{}) (*HTMLGenerator, error) {
	// проверка прав работы в FS
	if err := os.WriteFile(dirTpl+"/test.txt", []byte("test"), 0o600); err != nil {
		return nil, err
	}
	if err := os.Remove(dirTpl + "/test.txt"); err != nil {
		return nil, err
	}

	tplGen := HTMLGenerator{
		functions: f,
		tplStore:  map[string]*template.Template{},
	}

	// первичная индексация и компиляция шаблонов
	err := filepath.Walk(dirTpl, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// каталоги исключения
		if _, ok := dirExcluded[info.Name()]; ok {
			return filepath.SkipDir
		}
		// проверка допустимых расширений
		if _, ok := fileParseExt[filepath.Ext(path)]; !ok || info.IsDir() {
			return nil
		}
		if err := tplGen.tplParse(dirTpl, path, f); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &tplGen, tplGen.watch(dirTpl)
}

// мониторинг, индексация и компиляция шаблонов при изменении FS
func (s *HTMLGenerator) watch(dirTpl string) error {
	t_watch := watcher.New()
	t_watch.IgnoreHiddenFiles(true)
	t_watch.FilterOps(watcher.Create, watcher.Write, watcher.Rename, watcher.Move, watcher.Remove)
	if err := t_watch.AddRecursive(dirTpl); err != nil {
		return err
	}

	go func() {
		for {
		start:
			fi := <-t_watch.Event
			// проверка допустимых расширений
			if _, ok := fileParseExt[filepath.Ext(fi.Path)]; !ok || fi.IsDir() || !fi.Mode().IsRegular() {
				continue
			}
			// каталоги исключения
			for i := range dirExcluded {
				if strings.Contains(fi.Path, i) {
					goto start
				}
			}
			switch fi.Op {
			default:
				continue
			case watcher.Create:
				_ = s.tplParse(dirTpl, fi.Path, s.functions)
			case watcher.Write:
				_ = s.tplParse(dirTpl, fi.Path, s.functions)
			case watcher.Rename:
				s.tplRemove(dirTpl, fi.OldPath)
				_ = s.tplParse(dirTpl, fi.Path, s.functions)
			case watcher.Move:
				s.tplRemove(dirTpl, fi.OldPath)
				_ = s.tplParse(dirTpl, fi.Path, s.functions)
			case watcher.Remove:
				s.tplRemove(dirTpl, fi.OldPath)
			}
		}
	}()
	go func() {
		_ = t_watch.Start(time.Second * 10)
	}()
	return nil
}

func (s *HTMLGenerator) tplRemove(dir, path string) {
	_ = os.Remove(path)
	s.mu.Lock()
	delete(s.tplStore, strings.ReplaceAll(path, dir+"/", ""))
	s.mu.Unlock()
}

func (s *HTMLGenerator) tplParse(dir, path string, functions map[string]interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	index := strings.ReplaceAll(path, dir+"/", "")
	tpl, err := template.New(index).Funcs(functions).Parse(string(data))
	if err != nil {
		return err
	}

	s.mu.Lock()
	s.tplStore[index] = tpl
	s.mu.Unlock()
	return nil
}

// GetStorageIndex индексы подготовленных шаблонов
//
// Sample from storage 'www':
// index.html
// page/page.html
func (s *HTMLGenerator) GetStorageIndex() []string {
	s.mu.RLock()
	t := make([]string, 0, len(s.tplStore))
	for i := range s.tplStore {
		t = append(t, i)
	}
	s.mu.RUnlock()
	return t
}

// ExecuteStorage сборка контента из подготовленного шаблона
func (s *HTMLGenerator) ExecuteStorage(viewIndex string, variables map[string]interface{}) (ret bytes.Buffer, err error) {
	s.mu.RLock()
	tpl, ok := s.tplStore[viewIndex]
	s.mu.RUnlock()
	if !ok {
		return ret, errors.New("not found tpl: " + viewIndex)
	}
	err = tpl.Execute(&ret, variables)
	return
}

// ExecuteFile компиляция html шаблона из указанного файла и сборка контента
func ExecuteFile(viewPath string, funcs, variables map[string]interface{}) (ret bytes.Buffer, err error) {
	if _, err = os.Stat(viewPath); err != nil {
		return ret, err
	}

	var tpl *template.Template

	/*if funcs == nil {
		funcs = Functions
	}*/

	if tpl, err = template.New(filepath.Base(viewPath)).Funcs(funcs).ParseFiles(viewPath); err != nil {
		return ret, err
	}

	err = tpl.Execute(&ret, variables)
	if err != nil {
		return ret, err
	}

	return
}

// ExecuteString компиляция html шаблона переданного в строке и сборка контента
func ExecuteString(view string, funcs, variables map[string]interface{}) (ret bytes.Buffer, err error) {
	const nameTpl = "default"
	var tpl *template.Template

	/*if funcs == nil {
		funcs = functions
	}*/

	if tpl, err = template.New(nameTpl).Funcs(funcs).Parse(view); err != nil {
		return ret, err
	}

	err = tpl.Execute(&ret, variables)
	if err != nil {
		return ret, err
	}

	return
}
