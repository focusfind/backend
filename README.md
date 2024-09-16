# focusfind backend

## API Endpoints

### Spot Operations

- GET /api/spots: Get a list of all spots within 1 mile radius.
- GET /api/spots/spotId: Get details of a specific spot.
- POST /api/spots: Add a new spot 


### Other

- PUT /api/spots/:spotId: Update details of a specific spot 
- DELETE /api/spots/:spotId: Delete a spot (requires authentication and ownership).


#### Spot Reviews

- GET /api/spots/:spotId/reviews: Get reviews for a specific spot.
- POST /api/spots/:spotId/reviews: Add a review for a spot (requires authentication).
- PUT /api/spots/:spotId/reviews/:reviewId: Update a review (requires authentication and ownership).
- DELETE /api/spots/:spotId/reviews/:reviewId: Delete a review (requires authentication and ownership).
- GET /api/spots?type=study_spot: Filter spots based on type.
