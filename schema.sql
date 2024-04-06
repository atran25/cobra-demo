CREATE TABLE "characters" (
    "character_id" text PRIMARY KEY,
    "server_id" text NOT NULL,
    "character_name" text NOT NULL,
    "job_id" text NOT NULL,
    "job_grow_id" text NOT NULL,
    "rank" int ,
    PRIMARY KEY ("id")
);