
/* Here we are using anonymous interface within type assertion to check if type x has a String method */
if _, ok := x.(interface {
	String() string
}); !ok {
	fmt.Printf("x can't String()")
}