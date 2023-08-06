package dish

// Delete: Delete single dish by ID
func (c DishController) Delete(id int) error {
	err := c.dishRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
