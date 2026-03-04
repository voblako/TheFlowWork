-- 1. Permissions (Права доступа)
CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    position VARCHAR(255) NOT NULL, -- должность
    work_access BOOLEAN NOT NULL, -- доступ к работе
    planning_access BOOLEAN NOT NULL, -- доступ к планированию (optional in source)
    issue_creation_access BOOLEAN NOT NULL, -- доступ к созданию проблем
    supplies_access BOOLEAN NOT NULL, -- доступ к расходникам
    fault_template_access BOOLEAN NOT NULL, -- доступ к шаблонам неполадок
    work_history_access BOOLEAN NOT NULL, -- доступ к истории работ
    workshop_editing_access BOOLEAN NOT NULL -- доступ к редактированию цехов
);

-- 2. Workshops (Цеха)
CREATE TABLE workshops (
    id SERIAL PRIMARY KEY,
    address TEXT NOT NULL, -- адрес
    type VARCHAR(255) NOT NULL, -- тип цеха сборочный, разгрузочный, ремонтный...
    manager_id INTEGER -- ответсвенный
);

-- 3. Equipment (Оборудование)
CREATE TABLE equipment (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    photo TEXT NOT NULL,
    workshop_id INTEGER NOT NULL, -- идентификатор цеха
    start_date DATE NOT NULL, -- дата начала эксплуатации
    last_maintenance_date DATE NOT NULL, -- дата последнего ТО
    CONSTRAINT fk_equipment_workshop FOREIGN KEY (workshop_id) REFERENCES workshops(id)
);

-- 4. Users (Пользователи)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL, -- имя
    password_hash TEXT NOT NULL, -- хэш пароль
    email TEXT NOT NULL,
    last_name TEXT NOT NULL, -- фамилия
    middle_name TEXT, -- отчество (optional)
    position INTEGER NOT NULL, --ссылка на право доступа (permissions id)
    workshop_id INTEGER NOT NULL, -- идентификатор цеха
    CONSTRAINT fk_users_workshop FOREIGN KEY (workshop_id) REFERENCES workshops(id),
    CONSTRAINT fk_users_permissions FOREIGN KEY (position) REFERENCES permissions(id)
);

-- Добавляем внешний ключ для manager_id в таблице workshops после создания таблицы users.
ALTER TABLE workshops 
ADD CONSTRAINT fk_workshop_manager FOREIGN KEY (manager_id) REFERENCES users(id);

-- 5. Requests/Tickets (Заявка)
CREATE TABLE requests (
    id SERIAL PRIMARY KEY,
    equipment_id INTEGER NOT NULL, -- идентификатор оборудования
    start_date DATE NOT NULL, -- сообщение о проблеме
    end_date DATE, -- работы выполнены (optional)
    problem_description TEXT NOT NULL, -- описание неполадки
    photo TEXT, -- optional
    CONSTRAINT fk_request_equipment FOREIGN KEY (equipment_id) REFERENCES equipment(id)
);

-- 6. Faults (Неполадки)
CREATE TABLE faults (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    photo TEXT, -- optional
    executor_id INTEGER NOT NULL, -- идентификатор пользователя
    created_at TIMESTAMP NOT NULL, -- дата создания
    closed_at TIMESTAMP, -- дата закрытия (optional)
    CONSTRAINT fk_fault_executor FOREIGN KEY (executor_id) REFERENCES users(id)
);

-- 7. Fault Solution Templates (Шаблоны решения неполадок)
CREATE TABLE fault_solution_templates (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    equipment_ids INTEGER[] NOT NULL, -- массив идентификаторов оборудования, к которому применим шаблон
    solution TEXT NOT NULL
);

-- 8. Actions (Действия)
CREATE TABLE actions (
    id SERIAL PRIMARY KEY,
    equipment_id INTEGER NOT NULL,
    comment TEXT NOT NULL,
    executor_id INTEGER NOT NULL, -- исполнитель (user id)
    media TEXT, -- ссылки на внешние хранилище фото/видео (optional)
    time TIMESTAMP NOT NULL, -- время
    CONSTRAINT fk_action_executor FOREIGN KEY (executor_id) REFERENCES users(id),
    CONSTRAINT fk_action_equipment FOREIGN KEY (equipment_id) REFERENCES equipment(id)
);

-- 9. Supplies (Расходники)
CREATE TABLE supplies (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    workshop INTEGER NOT NULL, -- цех
    description TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    CONSTRAINT fk_supplies_location FOREIGN KEY (workshop) REFERENCES workshops(id)
);

-- 10. Schedule (Расписание)
CREATE TABLE schedule (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    priority INTEGER NOT NULL,
    scheduled_time TIMESTAMP NOT NULL, -- время
    executor_id INTEGER NOT NULL, -- исполнитель (user id)
    equipment_id INTEGER NOT NULL, -- оборудование
    completed_at TIMESTAMP, -- когда исполнено (в начале пустое, заполняется после исполнении)
    CONSTRAINT fk_schedule_executor FOREIGN KEY (executor_id) REFERENCES users(id),
    CONSTRAINT fk_schedule_equipment FOREIGN KEY (equipment_id) REFERENCES equipment(id)
);