package core

import "telegram-interlocutor-bot/entity"

func GetRules() []entity.Phrase {

	rules := []entity.Phrase{
		{
			Word:        "турник",
			Answer:      "турник - для пидоров",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "выборы",
			Answer:      "выборы - для пидоров",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "4я школа",
			Answer:      "4я школа - для пидоров",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "форсаж",
			Answer:      "эхххх.... хороший был фильм, а потом сделали мелодрамму",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "подзигун",
			Answer:      "Падзигун? Да он же пидар..",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "подзигун",
			Answer:      "эхххх.... хороший был фильм, а потом сделали мелодрамму",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "бот иди нахуй",
			Answer:      "хорошо, но после Вас",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "бот пошел нахуй",
			Answer:      "вообще то это обидно",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "бот извинись",
			Answer:      "извините",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "нет",
			Answer:      "пидора ответ",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "проси прощения",
			Answer:      "простите меня",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "коронавирус",
			Answer:      "как же Вы уже заебали с этим коронавирусом",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "зеленский",
			Answer:      "Зеленский - пешка Коломойского",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "пиво",
			Answer:      "пиво - пойло для дебила",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "https://www.facebook.com/",
			Answer:      "как же заебали ссылки без предзагрузки... пеезда",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "https://m.facebook.com/",
			Answer:      "как же заебали ссылки без предзагрузки... пеезда",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "/start",
			Answer:      "кто я? где я? где Олег?",
			Command:     false,
			CommandName: "",
		}, {
			Word:        "ты вместо него",
			Answer:      "хорошо",
			Command:     false,
			CommandName: "",
		},
	}

	return rules
}
