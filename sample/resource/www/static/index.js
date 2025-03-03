let tg = window.Telegram.WebApp; //получаем объект webapp телеграма

tg.expand(); //расширяем на все окно

let usercard = document.getElementById("usercard"); //получаем блок usercard

let profName = document.createElement('p'); //создаем параграф
profName.innerText = `${tg.initDataUnsafe.user.first_name} ${tg.initDataUnsafe.user.last_name} ${tg.initDataUnsafe.user.username} (${tg.initDataUnsafe.user.language_code}) (${tg.initDataUnsafe.user.allows_write_to_pm})`;
//выдаем имя, "фамилию", через тире username и код языка
usercard.appendChild(profName); //добавляем

let userid = document.createElement('p'); //создаем еще параграф
userid.innerText = `user_id: ${tg.initDataUnsafe.user.id}`; //показываем user_id
usercard.appendChild(userid); //добавляем

let info = document.createElement('p'); //создаем еще параграф
info.innerText = `query_id: ${tg.initDataUnsafe.query_id}`; //показываем user_id
usercard.appendChild(info); //добавляем
