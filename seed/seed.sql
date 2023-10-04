CREATE TABLE if not exists users (
                                     id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                                     "first_name" varchar(55) NOT NULL,
                                     "last_name" varchar(55) NOT NULL,
                                     "email" varchar(55) NOT NULL unique,
                                     "password" varchar(255) NOT NULL,
                                     "access" int4 DEFAULT 0,
                                     "inserted_at" timestamp NOT NULL DEFAULT now(),
                                     "updated_at" timestamp
);

CREATE TABLE IF NOT EXISTS place_comment (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    place_id uuid NOT NULL,
    comment text not null,
    "inserted_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS places (
                                      id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                                      name varchar(255) NOT NULL,
                                      image_url varchar(255) ,
                                      description text,
                                      location varchar(255) ,
                                      instagram varchar(55) ,
                                      ig_url varchar(255),
                                      contact varchar(65),
                                      "inserted_at" timestamp NOT NULL DEFAULT now(),
                                      "updated_at" timestamp
)