
## Описание проекта

Backend-сервис предоставляет возможность пользователям выставлять на продажу квартиры. Сервис предоставляет функционал для создания, просмотра и модерации объявлений, а также регистрации и авторизации пользователей.

### Функционал сервиса: 

1. **Авторизация пользователей:**
	Регистрация и авторизация пользователей по почте и паролю:
	-  При регистрации используется endpoint /register. В базе создаётся и сохраняется новый пользователь желаемого типа: обычный пользователь (client) или модератор (moderator).
    - У созданного пользователя появляется токен endpoint /login. При успешной авторизации по почте и паролю возвращается токен для пользователя с соответствующим ему уровнем доступа.
2. **Создание дома:**
    1. Только модератор имеет возможность создать дом используя endpoint /house/create. В случае успешного запроса возвращается полная информация о созданном доме.
3. **Создание квартиры:**
    1. Создать квартиру может любой пользователь, используя endpoint /flat/create. При успешном запросе возвращается полная информация о квартире.
    2. Если жильё успешно создано через endpoint /flat/create, то объявление получает статус модерации created.
    3. У дома, в котором создали новую квартиру, обновляется дата последнего добавления жилья. 
4. **Модерация квартиры:**
    1. Статус модерации квартиры может принимать одно из четырёх значений: created, approved, declined, on moderation.
    2. Только модератор может изменить статус модерации квартиры с помощью endpoint /flat/update. При успешном запросе возвращается полная информация об обновленной квартире.
5. **Получение списка квартир по номеру дома:**
    1. Используя endpoint /house/{id}, обычный пользователь и модератор могут получить список квартир по номеру дома. Только обычный пользователь увидит все квартиры со статусом модерации approved, а модератор — жильё с любым статусом модерации.

### Общие вводные

У сущности **«Дом»** есть:  

- уникальный номер дома
- адрес
- год постройки
- застройщик (у 50% домов)
- дата создания дома в базе
- дата последнего добавления новой квартиры дома

У сущности **«Квартира»** есть:

- номер квартиры
- цена (целое число)
- количество комнат

**Связи между сущностями:**

1. Каждая квартира может иметь только одно соответствие с домом (один к одному).
2. Номер дома служит уникальным идентификатором самого дома.
3. Номер квартиры не является уникальным идентификатором. Например, квартира №1 может находиться как в доме №1, так и в доме №2, и в этом случае это будут разные квартиры.

Список квартир в доме — ключевая функция, которой пользуются: 

- Модераторы — получают полный список всех объявлений в доме вне зависимости от статуса модерации.
- Пользователи — получают список только прошедших модерацию объявлений.  Важно гарантировать быстрый отклик endpoint для пользователей. 

## Условия и требования

1. Использовать этот [API](https://github.com/mrgoshha/house_service/blob/master/api/api.yaml).
2. Разработать интеграционные и модульные тесты для сценариев получения списка квартир и процесса публикации новой квартиры.
3. Квартира может не пройти модерацию. В таком случае её видят только модераторы. 
4. Работать с сервисом могут несколько модераторов. При этом конкретную квартиру может проверять только один модератор. Перед началом работы нужно перевести квартиру в статус on moderate — тем самым запретив брать её на проверку другим модераторам. В конце квартиру переводят в статус approved или declined.
5. Настроить логгер
6.  Для деплоя зависимостей и самого сервиса использовать Docker или Docker & DockerCompose.
