import java.io.*;
import java.util.*;
public class ContiguousSubArray {
    public static void main(String args[] ) throws Exception {

        Scanner reader = new Scanner(System.in);  // Reading from System.in
        System.out.println("Enter a number: ");
        int n = reader.nextInt();
        int arr[]=new int[n];
        for(int i=0; i<n; i++) {
            arr[i] = reader.nextInt();
        }

        System.out.println(MaxSubArraySum(arr));
    }

    static int MaxSubArraySum(int arr[]) {
        int size = arr.length;
        int max_sum = arr[0];
        int sum = arr[0];
        for(int i=1; i<size; i++) {
            sum += arr[i];
            if(sum < 0)
                sum = 0;
            if(max_sum < sum)
                max_sum = sum;
        }
        return max_sum;
    }
}
