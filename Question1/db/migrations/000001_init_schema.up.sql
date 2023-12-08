CREATE TABLE "users" (
    "id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
    "name" VARCHAR NOT NULL,
    "phone_number" VARCHAR NOT NULL,
    "otp" VARCHAR NULL,
    "otp_expiration_time" TIMESTAMP(3) NULL,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "users_phone_number_key" ON "users"("phone_number");