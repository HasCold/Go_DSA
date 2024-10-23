// When you run three separate loops in a function, each loop executes independently and sequentially, not nested inside each other. This is key to understanding why the overall time complexity is still \( O(n) \) and not something more complicated.

// Here’s how it works:

// ### Example of Three Separate Loops:
// ```go
// func exampleFunction(arr []int) {
//     // First loop runs n times
//     for i := 0; i < len(arr); i++ {
//         // Some O(1) operation
//     }

//     // Second loop runs n times
//     for i := 0; i < len(arr); i++ {
//         // Some O(1) operation
//     }

//     // Third loop runs n times
//     for i := 0; i < len(arr); i++ {
//         // Some O(1) operation
//     }
// }
// ```

// ### Breaking it down:
// - **First loop:** Runs \( n \) times.
// - **Second loop:** Also runs \( n \) times.
// - **Third loop:** Again, runs \( n \) times.

// So the total number of operations is \( n + n + n = 3n \).

// ### Why is it \( O(n) \), not \( O(3n) \)?
// - **Big O notation** focuses on the *asymptotic behavior* of the function, which means how the function behaves as the input size \( n \) grows very large.
// - Constants like 3 don’t affect the asymptotic behavior. In Big O, we ignore constants because they do not fundamentally change the growth rate as \( n \) increases. Therefore, \( O(3n) \) simplifies to \( O(n) \).

// ### Intuition:
// - \( O(n) \) means that the execution time grows linearly with the size of the input. Whether it’s \( n \) operations or \( 3n \), both represent linear growth, which is why we drop the constant factor and keep it as \( O(n) \).

// Thus, even though there are 3 loops each running \( n \) times, the overall time complexity is still \( O(n) \) because constant factors are ignored in Big O notation.

package main

func main() {

}
