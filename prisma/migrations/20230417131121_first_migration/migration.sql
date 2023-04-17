-- CreateTable
CREATE TABLE `BusTravel` (
    `id` INTEGER NOT NULL AUTO_INCREMENT,
    `origin` VARCHAR(191) NOT NULL,
    `destination` VARCHAR(191) NULL,
    `departureTime` DATETIME(3) NOT NULL,
    `arrivalTime` DATETIME(3) NOT NULL,
    `price` DOUBLE NOT NULL,
    `hasWiFi` BOOLEAN NULL,
    `hasBathroom` BOOLEAN NULL,
    `hasAirConditioner` BOOLEAN NULL,

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
