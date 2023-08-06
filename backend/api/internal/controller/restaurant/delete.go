package restaurant

// Delete: Delete single restaurant by ID
func (c RestaurantController) Delete(id int) error {
	err := c.restaurantRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
