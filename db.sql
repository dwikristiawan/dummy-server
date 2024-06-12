CREATE TABLE workspaces (
                            id VARCHAR PRIMARY KEY,
                            name VARCHAR NOT NULL,
                            reference_id VARCHAR,
                            created_at TIMESTAMP WITH TIME ZONE,
                            updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE members (
                         id VARCHAR PRIMARY KEY,
                         workspace_id VARCHAR NOT NULL,
                         user_id VARCHAR NOT NULL,
                         access VARCHAR NOT NULL,
                         is_active BOOLEAN NOT NULL,
                         created_at TIMESTAMP WITH TIME ZONE,
                         updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE collections (
    id VARCHAR PRIMARY KEY,
    workspace_id VARCHAR NOT NULL,
    reference_id VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE

);
CREATE TABLE childrens (
    id VARCHAR PRIMARY KEY,
    collection_id VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    perent VARCHAR NOT NULL,
    reference_id VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE

);

CREATE TABLE mockdatas (
                           id VARCHAR PRIMARY KEY,
                           children_id VARCHAR NOT NULL,
                           collection_id VARCHAR NOT NULL,
                           request_method VARCHAR NOT NULL,
                           path VARCHAR,
                           request_header JSON,
                           response_header JSON,
                           request_body VARCHAR,
                           response_body VARCHAR,
                           response_code INT,
                           reference_id VARCHAR NOT NULL,
                           created_at TIMESTAMP WITH TIME ZONE NOT NULL ,
                           updated_at TIMESTAMP WITH TIME ZONE
);