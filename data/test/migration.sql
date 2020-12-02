CREATE TABLE "public"."users" (
    "id" uuid NOT NULL,
    "name" character varying(120) NOT NULL,
    "username" character varying(60) NOT NULL,
    CONSTRAINT "users_id" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "users_name" ON "public"."users" USING btree ("name");

CREATE INDEX "users_username" ON "public"."users" USING btree ("username");

CREATE TABLE "public"."relevance_1" (
    "id" uuid NOT NULL,
    CONSTRAINT "relevance_1_id_fkey" FOREIGN KEY (id) REFERENCES users(id) NOT DEFERRABLE
) WITH (oids = false);

CREATE TABLE "public"."relevance_2" (
    "id" uuid NOT NULL,
    CONSTRAINT "relevance_2_id_fkey" FOREIGN KEY (id) REFERENCES users(id) NOT DEFERRABLE
) WITH (oids = false);
