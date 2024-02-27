package test

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/shopspring/decimal"

	"gitlab.tn.ru/golang/kit/tpl"
)

// TestGeneratePdf
// apt-get -y install wkhtmltopdf
func TestGeneratePdf(t *testing.T) {
	// template html
	f := map[string]interface{}{
		"TplTest": func(name string) string {
			return "<H1>" + name + "</H1>"
		},
	}
	goods := Goods{
		{ID: 37, Name: "Item 10", Price: decimal.NewFromFloat(23.76)},
		{ID: 49, Name: "Item 2", Price: decimal.NewFromFloat(87.42)},
		{ID: 54, Name: "Item 30", Price: decimal.NewFromFloat(38.23)},
	}

	pic, err := tpl.GetImgBase64File("6665119_5_l.jpeg")
	if err != nil {
		t.Fatal(err)
	}

	variable := map[string]interface{}{
		"Title":    "TestTplStorage",
		"ImgLoser": pic,
		"Goods":    goods,
	}

	// generate pdf
	genPdf := tpl.NewPDFGenerator()

	// 1 file
	bufTpl, err := tpl.ExecuteString(testTpl, f, variable)
	if err != nil {
		t.Fatal(err)
	}

	if err := genPdf.GenerateFile(&bufTpl, "TestPdfFile.pdf"); err != nil {
		t.Fatal(err)
	}

	// 2 buffer
	bufTpl, err = tpl.ExecuteString(testTpl, f, variable)
	if err != nil {
		t.Fatal(err)
	}

	obj := &bytes.Buffer{}
	if err := genPdf.Generate(&bufTpl, obj); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile("TestPdfBuffer.pdf", obj.Bytes(), 0o666); err != nil {
		t.Fatal(err)
	}

	// 3 transport
	bufTpl, err = tpl.ExecuteString(testTpl, f, variable)
	if err != nil {
		t.Fatal(err)
	}

	// формирующий сервис
	data, err := genPdf.GenerateToJson(&bufTpl)
	if err != nil {
		t.Fatal(err)
	}

	// сохраняющий сервис
	var outPDF bytes.Buffer
	if err := genPdf.GenerateFromJson(data, &outPDF); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile("TestPdfTransport.pdf", outPDF.Bytes(), 0o600); err != nil {
		t.Fatal(err)
	}
}

func TestGeneratePdfCPUAddDelAttach(t *testing.T) {
	filename := "TestPdfCPUAdd.pdf"
	// template html
	param := map[string]interface{}{
		"Attorney_date": time.Now().Format("2006-01-02"),
		"Entity":        "ООО Ромашка",
		"Ogrn":          "12345",
		"Inn":           "123456",
		"Address":       "адрес",
		"Number":        "998877",
		"Attorney_fio":  fmt.Sprintf("%s %s %s", "Иванов", "Иван", "Иванович"),
	}
	bufTpl, err := tpl.ExecuteString(RetroBonusSimpleEdo, nil, param)
	if err != nil {
		t.Fatal(err)
	}

	genPdf := tpl.NewPDFGenerator()

	err = genPdf.GenerateFile(&bufTpl, filename)
	if err != nil {
		t.Fatal(err)
	}

	err = genPdf.AddAttachment(filename, []string{"6086793_4_l.jpg", "6665119_5_l.jpeg"})
	if err != nil {
		t.Fatal(err)
	}

	// err = genPdf.DelAttachment(filename, []string{"6086793_4_l.JPG"})
	// if err != nil {
	// 	t.Fatal(err)
	// }
}

const RetroBonusSimpleEdo = `
<!DOCTYPE html>
<html lang="ru">
<head>
  <meta charset="UTF-8"/>
  <title>First</title>
  <style>
    .document, .document * {
      margin: 0;
      border: none;
      outline: none;
      font-size: 16px;
      text-align: justify;
    }

    .document {
      margin: 0 auto;
    }

    .document__header {
      font-weight: bold;
      text-align: center;
      margin-top: 45px;
      font-size: 21px;
    }

    .document__subheader {
      font-size: 18px;
      font-weight: bold;
      text-align: center;
      display: block;
    }

    .document__main {
      padding: 20px;
    }

    .document__main-text {
      text-align: justify;
    }

    .document__main-text_left {
      text-align: end;
    }

    .document__main-list-item {
      margin: 20px 0;
    }

    .document__main-list-item > .document__subheader {
      margin-bottom: 20px;
    }

    ol {
      padding: 0;
      counter-reset: item;
    }

    .document li {
      padding-left: 40px;
      position: relative;
      display: block;
    }

    li li::before {
      font-size: 16px;
      font-weight: normal;
    }

    li::before {
      left: 0;
      position: absolute;
      font-size: 18px;
      font-weight: bold;
      content: counters(item, ".") " ";
      counter-increment: item;
    }

    h3 {
      margin-bottom: 20px;
    }

    .document .document__main-list-item.no-padding {
      padding-left: 0;
    }
  </style>
</head>
<body class="document">
<main class="document__main">
    <h1 class="document__header">Соглашение на осуществление документооборота в электронном виде для подписания документов, связанных с участием в Программах Продавца о вознаграждениях
        <br/>
        ОФЕРТА
    </h1>
    <p class="document__main-text document__main-text_left">{{.Attorney_date}}</p>
    <p class="document__main-text">Данный документ является официальным предложением <strong>ООО &laquo;РОМАШКА&raquo;</strong> (далее по тексту – «<strong>Продавец</strong>») в адрес:</p>
    <p class="document__main-text"><strong>{{.Entity}}</strong>, именуемое в дальнейшем «<strong>Покупатель</strong>»,
        заключить Соглашение на осуществление документооборота в электронном виде для принятия участия в Программах Продавца
        о вознаграждениях (далее – Соглашение), на указанных ниже условиях. Соглашение соответствии со статьей 435 ГК РФ
        является офертой и заключается путем акцепта оферты Покупателем. Письменная форма договора считается соблюденной в
        силу ст. 434 ГК РФ.</p>
    <ol>
        <li class="document__main-list-item no-padding">
            <h2 class="document__subheader">ТЕРМИНЫ, ОПРЕДЕЛЕНИЯ И СОКРАЩЕНИЯ</h2>
            <p class="document__main-text">Для целей настоящего документа в нем используются следующие термины, определения и
                сокращения, а именно:</p>
            <p class="document__main-text"><strong>Продавец</strong> – ООО &laquo;РОМАШКА&raquo; (ИНН
                7702521529 ОГРН 1048836256694), адрес местонахождения: 111110, г. Москва, ул. Ленина, д.44, стр.4, эт.4,
                пом. 4, к.14.</p>
            <p class="document__main-text"><strong>Акцепт Соглашения</strong> – полное принятие Соглашения путем осуществления
                действий, предусмотренных настоящем Соглашением с использованием простой электронной подписи;</p>
            <p class="document__main-text"><strong>Программы Продавца о вознаграждениях</strong> – это стимулирующие
                программы, проводимые Продавцом для Покупателей (далее – Программы Продавца);</p>
            <p class="document__main-text"><strong>Личный кабинет (ЛК)</strong> интерфейс в сети Интернет по адресу:
                https://b2b.tn.ru, предоставляющий возможность Покупателю принимать условия участия в Программах Продавца и получать
                актуальную информацию по ним;</p>
            <p class="document__main-text"><strong>Простая Электронная подпись (ПЭП)</strong> (в соответствии с Федеральным
                законом «Об электронной подписи» от 06.04.2011г. N 63-ФЗ) - Электронная подпись, которая посредством
                использования кодов, паролей или иных средств подтверждает факт формирования электронной подписи определенным
                лицом.</p>
            <p class="document__main-text"><strong>Электронный документ</strong> – документированная информация,
                представленная в электронной форме, т.е. в виде, пригодном для восприятия человека с использованием электронных
                вычислительных машин, а также для передачи по информационно-телекоммуникационным сетям или обработки в
                информационных системах;</p>
            <p class="document__main-text"><strong>Электронный документооборот (ЭДО)</strong> – способ взаимодействия Сторон
                по обмену Электронными документами, подписанными ПЭП с целью участия в Программах Продавца.</p>
        </li>
        <li class="document__main-list-item">
            <h2 class="document__subheader">ПРЕДМЕТ СОГЛАШЕНИЯ</h2>
            <ol>
                <li>Настоящим Соглашением Стороны определяют условия и порядок организации обмена юридически значимыми
                    электронными документами в рамках ЭДО с использованием ПЭП в качестве аналога собственноручной подписи
                    уполномоченного представителя Покупателя и печати организации по вопросам принятия и исполнения условий
                    Программ Продавца.
                </li>
                <li>Получение документов в рамках Программ Продавца в электронном виде и подписанных ПЭП в порядке,
                    установленным настоящим Соглашением, эквивалентно получению документов на бумажном носителе и является
                    необходимым и достаточным условием, позволяющим установить, что ЭД исходит от Стороны, его направившей.
                </li>
                <li>Документы в рамках Программ Продавца, направленные одной Стороной другой Стороне в электронном виде, и
                    подписанные Покупателем с использование ПЭП, считаются заключенными между Сторонами в простой письменной
                    форме, согласно п. 2 ст. 434 ГК РФ.
                </li>
            </ol>
        </li>
        <li class="document__main-list-item">
            <h2 class="document__subheader">ПОДПИСАНИЕ ДОКУМЕНТОВ</h2>
            <ol>
                <li>Принятием условий настоящего Соглашения и Программ Продавца, а также их подписанием является ввод
                    Покупателем разового цифрового кода (пароля), полученного на номер мобильного телефона уполномоченного
                    представителя Покупателя, посредством sms от Продавца, который вносится в соответствующее поле в ЛК.
                </li>
            </ol>
        </li>
        <li class="document__main-list-item">
            <h2 class="document__subheader">ПРАВА И ОБЯЗАННОСТИ СТОРОН</h2>
            <ol>
                <li>Продавец обязуется направить на номер мобильного телефона уполномоченного представителя Покупателя,
                    указанный в предоставленной Покупателем доверенности, sms, содержащее пароль, необходимый для принятий условий
                    настоящего Соглашения или Программ Продавца. Переданный пароль является Электронной подписью Покупателя на
                    момент подписания Соглашения или Программ Продавца.
                </li>
                <li>Покупатель гарантирует допуск к ЭДО только уполномоченных сотрудников и несет ответственность за все
                    действия, совершенные неуполномоченными лицами. В каждом случае получения подписанного электронного документа
                    Продавец добросовестно исходит из того, что документ подписан от имени Покупателя надлежащим лицом,
                    действующим в пределах, имеющихся у него полномочий
                </li>
                <li>Покупатель обязан предотвращать несанкционированное использование третьими лицами соответствующего пароля от
                    его имени, полученного от Продавца, согласно настоящего Соглашения.
                </li>
                <li>После заключения настоящего Соглашения и в течение всего срока его действия Покупатель будет иметь право
                    подписывать (в порядке, предусмотренном п.3.1 настоящего Соглашения) своей ПЭП документы по вопросам принятия
                    и исполнения условий Программ Продавца. Стороны подтверждают, что любой документ, направленный Покупателю
                    Продавцом через ЛК, в случае его подписания ПЭП Покупателя, в соответствии с Соглашением, считается
                    подписанным непосредственно Покупателем.
                </li>
            </ol>
        </li>
        <li class="document__main-list-item">
            <h2 class="document__subheader">ПОРЯДОК РАЗРЕШЕНИЯ СПОРОВ</h2>
            <ol>
                <li>Стороны принимают необходимые меры к тому, чтобы любые спорные вопросы, разногласия либо претензии, которые
                    могут возникнуть по поводу настоящего Соглашения, были урегулированы, прежде всего, путем взаимных
                    переговоров. Срок направления ответа на претензию не должен превышать 30 (тридцати) календарных дней со дня ее
                    получения Стороной
                </li>
                <li>Все споры и разногласия, возникающие из настоящего Соглашения или в связи с ним, в том числе касающиеся его
                    выполнения, нарушения, прекращения или действительности, не урегулированные путем переговоров, подлежат
                    разрешению в соответствии с действующим законодательством Российской Федерации в Арбитражном суде г. Москвы.
                </li>
            </ol>
        </li>
        <li class="document__main-list-item">
            <h2 class="document__subheader">СРОК ДЕЙСТВИЯ СОГЛАШЕНИЯ, ПОРЯДОК ЕГО ИЗМЕНЕНИЯ И РАСТОРЖЕНИЯ</h2>
            <ol>
                <li>Настоящее Соглашение вступает в силу с даты принятия Покупателем условий и подписание с использованием ПЭП
                    настоящего Соглашения и считается заключенным на неопределенный срок.
                </li>
                <li>
                    Соглашение может быть расторгнуто:
                    <ol>
                        <li>По соглашению Сторон;</li>
                        <li>По инициативе любой Стороны. При условии уведомления другой стороны о расторжении не менее, чем за 3
                            (три) календарных дня до предполагаемой даты расторжения Соглашения;
                        </li>
                    </ol>
                </li>
                <li>Любые изменения и дополнения к настоящему Соглашению действительны только в том случае, если они совершены в
                    письменной форме, в том числе с использованием ЭДО, в этом случае они являются его неотъемлемой частью.
                </li>
            </ol>
        </li>
    </ol>
    <h3>РЕКВИЗИТЫ ПОКУПАТЕЛЯ:</h3>
    <p>
        {{.Entity}}
        <br/>
        ОГРН {{.Ogrn}}
        <br/>
        ИНН {{.Inn}}
        <br/>
        Адрес {{.Address}}
        <br/>
        <br/>
        Представить по доверенности № {{.Number}} от {{.Attorney_date}}
        <br/>
        ФИО {{.Attorney_fio}}
    </p>
</main>
</body>
</html>`
