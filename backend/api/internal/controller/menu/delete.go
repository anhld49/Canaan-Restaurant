package menu

// Delete: Delete single menu by ID
func (c MenuController) Delete(id int) error {
	err := c.menuRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
