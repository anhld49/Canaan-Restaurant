package order

// Delete: Delete single order by ID
func (c OrderController) Delete(id int) error {
	err := c.orderRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
