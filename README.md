1. All the API will be Paginated 
2. Created API should authenticated API Endpoints. 
3. Send a Notification Using Go-Routine-Channel-Worker Channel who has reported the same tigers. 
4. Migration Should be Done Initially When you are setting up your Databases. 
5. ReadMe Need to change depending upon the docker-compose files.


## API Endpoints 
1. Create a New User (user_name, email, password). 
2. Login : How the user-should be able to login now 
3. Create a New Tiger : Name, Date of birth, Last-Seen Time, Last Lat-long
4. Listing of all the Tigers : Sorted by the Time we have Last Time. 
5. Create a New Tiger with the sighting : lat,long, timestamp and also add the support for image uploaded we have. 
6. List all sightings of a tiger : Tiger kha-2 seens huwa hai list dena 

User: 
- id
- name 
- username
- email 
- phone_number 
- password 
- active 
- created_at 
- updated_at 


Login ke time :
- user_id 
- token 
- active 
- expiry_time 
- created_at 
- updated_at 

Tiger
- name 
- dob 
- created_at 
- updated_at 


TigerSight
- tiger_id 
- last_seen 
- lat
- long 
- image_url 
- active 
- created_at 
- updated_at 
