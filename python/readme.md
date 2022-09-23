# Create User Config
## building_id
Used to specify which building these users should be created for, must be a valid ID of a building in the database. Specifies the `building_id` field in users.

## config fields:
- **num_users**: `integer`  
The number of users that will be generated
- **password_override**: `string`  
If not null, all users created will have this password
- **email_domains**: `array of strings`  
A list of possible domains that will be chosen from at random
- **lowest_preferred_start_time**: `datetime in ISO8601 format`  
Specifies the earliest allowable times that generated users can come in
- **highest_preferred_end_time**: `datetime in ISO8601 format`  
Specifies the latest end time that can be assigned to generated users
- **preferred_time_step_minutes**: `integer`  
A number representing minutes, this step will be used to create possible start/end time bins, which will be chosen from for a user, example: If the earliest start time is 07:00 and the latest end time is 16:00, with a step of 120 minutes, the possible start end times will be [07:00, 09:00, 11:00, 13:00, 15:00]
- **work_from_home_probability**: `float in range [0, 1]`   
Specifies the probability that the work_from_home attribute of users will be true
- **team_probabilities**: `array containing floats, must sum to 1`  
This dictates the assignment of users to teams, the index of the array corresponds to the number of teams the user can be in e.g. index represents 0 teams, index 5 represents 5 teams. The numbers in the arrays represent the probability that the user will be in the number of teams that the index represents. Hence, the array [0.1, 0.3, 0.6] means a generated user will have a 10% chance of not being in a team, a 30% chance of being in 1 team, and a 60% chance of being in 2 teams
- **role_probabilities**: `array containing floats, must sum to 1`  
Identical to team_probabilities but works with roles instead of teams
- **office_days_options**: `array of integers`  
The office_days field in generated users will be set to one of these integers
- **office_days_probabilities**: `array of floats that must sum up to one, and the length must match that of office_days_options`  
Each probability corresponds to the office days option in the same position, and represents the probability that that option is chosen
- **no_preferred_desk_probability** `float in range [0, 1]`  
A probability determine the chance that a user has no preferred desk