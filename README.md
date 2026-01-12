# 📅 ScheduleBot 



<p align="center">
  <a href="#-english">English</a> • 
  <a href="#-русский">Русский</a>
</p>

---

## 🇺🇸 English 

### About
**ScheduleBot** is a Telegram bot designed to parse and deliver college schedules directly to your phone. It's built for speed and reliability, ensuring you can access your classes even with a poor internet connection.

### Architecture
The project is split into two components to bypass regional IP restrictions (the university website only responds to Russian IP addresses):



1.  **Client (Parser):** Written in **Go**, runs within Russia. It scrapes data from the college website and uploads it to **Google Drive**.
2.  **Server (Bot):** Written in **Go**, hosted in the Netherlands. It pulls data from Google Drive and serves it to users via Telegram.

> **Why this way?** This hybrid approach is more cost-effective than using proxies and ensures a stable connection between the server and the university's local resources.

**Current Bot:** [@sсhedule_ASU_BOT](https://t.me/schedule_ASU_BOT)

---

## 🇷🇺 Русский

### О проекте
**ScheduleBot** — это Telegram-бот, который парсит расписание с сайта колледжа и предоставляет его в удобном виде. Это гораздо быстрее и стабильнее, чем каждый раз открывать сайт, особенно при слабом мобильном интернете.

### Архитектура
Проект разделен на две части, чтобы обойти региональные блокировки (сайт университета принимает запросы только с российских IP-адресов):

1.  **Клиент (Парсер):** Написан на **Go**, запущен в РФ. Собирает данные с сайта и загружает их на **Google Drive**.
2.  **Сервер (Бот):** Написан на **Go**, сервер находится в Нидерландах. Бот забирает данные из облака и отправляет их пользователям.

> **Зачем это нужно?** Мой сервер находится в Нидерландах, а сайт вуза блокирует зарубежные IP. Использование Google Drive как посредника — это дешевле и надежнее, чем аренда прокси-серверов.

**Ссылка на бота:** [@sсhedule_ASU_BOT](https://t.me/schedule_ASU_BOT)

---

## 🛠 Tech Stack / Стек технологий
* **Language:** Go (Golang)
* **API:** IDK (now)
* **Storage:** Google Drive API (as a bridge)
