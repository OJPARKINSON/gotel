# Gotel

## Intro

Building a quick project to run through a system design book that explores building a hotel booking system for a large international company with 5,000 hotels and 1 million rooms. The main objective is to build a full stack application that takes reservations and manages rooms. Personally for this project, the FE will be out of scope at least for me to hand build, I will let the vibes take over. My main objective is to test my architecture implimentation skills and do more golang and relational dbs.

## Data model

- Hotel Service 
    - Hotel table - basic info (name, address, location)
    - Room table - room info (room type, floor, number, hotel_id, name, availability)

- Rate Service
    - Room Type Rates table - (Hotel_id, date, rate)

- Guest Service
    - Guest table - (first, last name, email)

- Reservation Service
    - Room Type Inventory table - (hotel_id, room_type_id, date, total_inventory, total_reserved)
    - Reservation table - (hotel_id, room_type_id, start, end_date, status, guest_id)


