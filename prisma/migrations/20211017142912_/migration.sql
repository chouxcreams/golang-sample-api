/*
  Warnings:

  - You are about to drop the column `finishedAt` on the `Task` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "Task" DROP COLUMN "finishedAt",
ADD COLUMN     "completed" BOOLEAN NOT NULL DEFAULT false;
