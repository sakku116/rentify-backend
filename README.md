# Rentify Backend (development)

## API Documentation
go to `/swagger/index.html` to see API documentation.

## CLI Commands
- `seed-superuser`
    - args (optional):
        args should be empty for default seed (superuser;superuser@gmail.com;superuser). Or args must be containing 3 strings for custom username, email and passwords

        example:
        ```
        go run . seed-superuser
        ```
        or
        ```
        go run .seed-superuser superuser1 superuser1@gmail.com superuserpass
        ```

## App Design
https://www.figma.com/file/29shW48NZeNPn1sOJS2tiq/Rent-House-UI?type=design&node-id=0-1&mode=design - designed by Dirga Alfian Komara
