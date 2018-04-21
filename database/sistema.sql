-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS "sistema"."admin";
CREATE TABLE "sistema"."admin" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "email" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ----------------------------
-- Table structure for device
-- ----------------------------
DROP TABLE IF EXISTS "sistema"."device";
CREATE TABLE "sistema"."device" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "description" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "chip_id" varchar(64) COLLATE "pg_catalog"."default",
  "state" int2 NOT NULL DEFAULT 1,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS "sistema"."tag";
CREATE TABLE "sistema"."tag" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "id_user" uuid NOT NULL,
  "tag" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "description" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "state" int2 NOT NULL DEFAULT 1,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ----------------------------
-- Table structure for tag_device
-- ----------------------------
DROP TABLE IF EXISTS "sistema"."tag_device";
CREATE TABLE "sistema"."tag_device" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "id_device" uuid NOT NULL,
  "id_tag" uuid NOT NULL,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "state" int2 DEFAULT 1
);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "sistema"."user";
CREATE TABLE "sistema"."user" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "email" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ----------------------------
-- Table structure for history
-- ----------------------------
DROP TABLE IF EXISTS "sistema"."history";
CREATE TABLE "sistema"."history" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "device" character varying COLLATE pg_catalog."default" NOT NULL,
    "user" character varying COLLATE pg_catalog."default" NOT NULL,
    "tag" character varying COLLATE pg_catalog."default" NOT NULL,
    "state" smallint NOT NULL,
    "created_at" timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO "sistema"."admin" VALUES ('dcee4808-de43-4688-a09f-b9dcd967946b', 'Admin', 'admin@mail.com', 'admin', '2018-03-20 12:34:10.103865', '2018-03-20 12:43:14.547984');

-- ----------------------------
-- Records of device
-- ----------------------------
INSERT INTO "sistema"."device" VALUES ('070ee2b6-cc32-4ae8-bea6-59a3196683e3', 'Acesso frontal', 'Acesso frontal', '9640801', 1, '2018-03-20 02:16:39.149383', '2018-03-28 00:39:43.028763');
INSERT INTO "sistema"."device" VALUES ('87d3f981-1c0d-4345-a421-fd4ed4a3b58d', 'Acesso lateral', 'Acesso lateral', '', 1, '2018-03-26 15:46:40.158735', '2018-03-28 00:39:58.02048');
INSERT INTO "sistema"."device" VALUES ('f9cc3786-cf44-42a0-8d47-66e14f163f19', 'Device 02', 'Descricao device 02', 'AABBCC', 1, '2018-03-28 15:15:44.31516', '2018-03-28 15:15:44.31516');

-- ----------------------------
-- Records of tag
-- ----------------------------
INSERT INTO "sistema"."tag" VALUES ('bb10af52-c9cd-4854-968a-9761f55dcd3c', 'e5be054d-37fe-499d-9e70-78dea410bd17', 'F046E487', 'Tag da casa', 1, '2018-03-17 00:15:36.43087', '2018-03-26 01:40:06.262128');
INSERT INTO "sistema"."tag" VALUES ('ac9a59ab-d016-47a4-8f5f-eee75f2d3777', '0136dc24-ff9c-4930-af27-70ba90d75b85', '0353B0AB', '', 1, '2018-03-24 13:55:36.512048', '2018-03-26 21:00:59.178521');

-- ----------------------------
-- Records of tag_device
-- ----------------------------
INSERT INTO "sistema"."tag_device" VALUES ('08ca9c5e-dcf1-46ce-bb7a-d6222ded2999', '070ee2b6-cc32-4ae8-bea6-59a3196683e3', 'bb10af52-c9cd-4854-968a-9761f55dcd3c', '2018-03-22 18:38:18.393961', '2018-03-27 18:33:03.562865', 1);
INSERT INTO "sistema"."tag_device" VALUES ('434fc2ad-8652-4d9e-8e9b-6cb72259d654', '070ee2b6-cc32-4ae8-bea6-59a3196683e3', 'ac9a59ab-d016-47a4-8f5f-eee75f2d3777', '2018-03-26 16:15:31.5425', '2018-03-27 19:36:08.474301', 0);

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO "sistema"."user" VALUES ('0136dc24-ff9c-4930-af27-70ba90d75b85', 'Roberto', 'roberto@mail.com', '2018-03-20 16:34:27.803336', '2018-03-21 17:41:27.577489');
INSERT INTO "sistema"."user" VALUES ('e5be054d-37fe-499d-9e70-78dea410bd17', 'Douglas', 'douglas.zuqueto@gmail.com', '2018-03-17 00:15:15.044437', '2018-03-22 16:10:27.69185');

-- ----------------------------
-- Function structure for trigger_update_timestamp
-- ----------------------------
DROP FUNCTION IF EXISTS "sistema"."trigger_update_timestamp"();
CREATE OR REPLACE FUNCTION "sistema"."trigger_update_timestamp"()
  RETURNS "pg_catalog"."trigger" AS $BODY$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Indexes structure for table admin
-- ----------------------------
CREATE INDEX "idx_admin_email" ON "sistema"."admin" USING btree (
  "email" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Indexes structure for table device
-- ----------------------------
CREATE UNIQUE INDEX "idx_chipid" ON "sistema"."device" USING btree (
  "chip_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Indexes structure for table tag
-- ----------------------------
CREATE UNIQUE INDEX "idx_tag" ON "sistema"."tag" USING btree (
  "tag" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Indexes structure for table user
-- ----------------------------
CREATE UNIQUE INDEX "email" ON "sistema"."user" USING btree (
  "email" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
-- ----------------------------
-- Triggers structure for table admin
-- ----------------------------
CREATE TRIGGER "update_admin_updated_at" BEFORE UPDATE ON "sistema"."admin"
FOR EACH ROW
EXECUTE PROCEDURE "sistema"."trigger_update_timestamp"();

-- ----------------------------
-- Triggers structure for table device
-- ----------------------------
CREATE TRIGGER "update_device_updated_at" BEFORE UPDATE ON "sistema"."device"
FOR EACH ROW
EXECUTE PROCEDURE "sistema"."trigger_update_timestamp"();

-- ----------------------------
-- Triggers structure for table tag
-- ----------------------------
CREATE TRIGGER "update_tag_updated_at" BEFORE UPDATE ON "sistema"."tag"
FOR EACH ROW
EXECUTE PROCEDURE "sistema"."trigger_update_timestamp"();

-- ----------------------------
-- Triggers structure for table tag_device
-- ----------------------------
CREATE TRIGGER "update_tag_device_updated_at" BEFORE UPDATE ON "sistema"."tag_device"
FOR EACH ROW
EXECUTE PROCEDURE "sistema"."trigger_update_timestamp"();

-- ----------------------------
-- Triggers structure for table user
-- ----------------------------
CREATE TRIGGER "update_user_updated_at" BEFORE UPDATE ON "sistema"."user"
FOR EACH ROW
EXECUTE PROCEDURE "sistema"."trigger_update_timestamp"();

-- ----------------------------
-- Primary Key structure for table admin
-- ----------------------------
ALTER TABLE "sistema"."admin" ADD CONSTRAINT "admin_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table device
-- ----------------------------
ALTER TABLE "sistema"."device" ADD CONSTRAINT "device_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table tag
-- ----------------------------
ALTER TABLE "sistema"."tag" ADD CONSTRAINT "tag_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table tag_device
-- ----------------------------
ALTER TABLE "sistema"."tag_device" ADD CONSTRAINT "tag_device_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "sistema"."user" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table history
-- ----------------------------
ALTER TABLE "sistema"."history" ADD CONSTRAINT "history_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table tag
-- ----------------------------
ALTER TABLE "sistema"."tag" ADD CONSTRAINT "fk_id_user" FOREIGN KEY ("id_user") REFERENCES "sistema"."user" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Keys structure for table tag_device
-- ----------------------------
ALTER TABLE "sistema"."tag_device" ADD CONSTRAINT "fk_id_device" FOREIGN KEY ("id_device") REFERENCES "sistema"."device" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;
ALTER TABLE "sistema"."tag_device" ADD CONSTRAINT "fk_id_tag" FOREIGN KEY ("id_tag") REFERENCES "sistema"."tag" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;
