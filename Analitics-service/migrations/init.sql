CREATE TABLE qa(
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    question TEXT NOT NULL UNIQUE,
    answer TEXT NOT NULL,
    request_count INT DEFAULT 0,
    creation_time TIMESTAMP DEFAULT NOW(),
    project_type_id BIGINT,
    FOREIGN KEY (project_type_id)  REFERENCES project_type (id)
);

CREATE TABLE unresolved_question(
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    base_question_id BIGINT DEFAULT 0,
    is_solved BOOLEAN DEFAULT false,
    client_id BIGINT NOT NULL DEFAULT 0,
    client_name VARCHAR(20) DEFAULT 'default_name' NOT NULL,
    creation_time TIMESTAMP DEFAULT NOW(),
    question TEXT DEFAULT 'empty question'::text,

    project_type_id BIGINT,
    FOREIGN KEY (project_type_id)  REFERENCES project_type (id),
    
    qa_id BIGINT,
    FOREIGN KEY (qa_id)  REFERENCES qa (id)
);

CREATE TABLE project_type(
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,  
    project_name VARCHAR(50) NOT NULL
);

