
fn main() {

    let int_c: i32 = 6;
    println!("{}", sum_to_n_a(int_c));
    println!("{}", sum_to_n_b(int_c));
    println!("{}", sum_to_n_c(int_c));

    // iterative approach
    fn sum_to_n_a(mut n: i32) -> i32{
        let mut sum = 0;
        while n != 0{
            sum+=n;
            n-=1;
        }
        return sum;
    }

    // recursive approach
    fn sum_to_n_b(n: i32) -> i32{
        // condition to exit once it reaches the end.
        if n==0{
            return 0;
        }
        return n+sum_to_n_b(n-1);
    }

    // gauss approach
    fn sum_to_n_c(n: i32) -> i32{
        return n * (n+1)/2;
    }

}