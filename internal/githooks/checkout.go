package githooks

// PostCheckout is invoked after a checkout is completed. It is typically used to perform
// post-checkout operations such as updating configurations or refreshing environment data.
const PostCheckout = "post-checkout"

// getCheckoutHooks returns a slice of strings representing the names of checkout-related Git hooks.
// These hooks include post-checkout, which is invoked after a checkout is completed.
func getCheckoutHooks() []string {
	return []string{
		PostCheckout,
	}
}
