/* Read input from STDIN. Print your output to STDOUT*/

import java.io.*;
import java.util.*;
public class MaximizeSales {
    static int[] generateMapMinMax(Map<Integer, Integer> vendorMap) {
        int min=-1, max=-1;
        for(Integer in : vendorMap.keySet()){
            if(min == -1 || max == -1) {
                min = in;
                max = in;
            }
            if(in < min)
                min = in;
            else if(in > max)
                max = in;
        }
        int []output = new int[2];
        output[0] = min;
        output[1] = max;
        return output;
    }

    static void printLowestPriceForItem(String[] s){
        Map<Integer, Integer> ham=new HashMap<Integer, Integer>();
        for(String in:s){
            String[] words = in.split(" ");
            int en = Integer.parseInt(words[1]);
            int price = Integer.parseInt(words[2]);
            for(int i= Integer.parseInt(words[0]); i<=en; i++){
                if(ham.get(i) == null) {
                    ham.put(i, price);
                } else {
                    ham.put(i, Math.min(ham.get(i), price));
                }
            }
        }
        int[] arr = generateMapMinMax(ham);
        int start = -1;
        for(int x = arr[0]; x <= arr[1]; x++){
            if(start == -1) {
                start = x;
            } else if (ham.get(x) != ham.get(start) ){
                System.out.printf("%s %s %s\n", start, x-1, ham.get(start));
                start = x;
            }
            if(x == arr[1]) {
                System.out.printf("%s %s %s\n", start, x, ham.get(start));
            }
        }
    }

    public static void main(String args[]) throws Exception {
        Scanner reader = new Scanner(System.in);
        reader.useDelimiter(System.getProperty("line.separator"));
        int n = Integer.parseInt(reader.nextLine());
        String arr[] = new String[n];
        for(int i=0; i<n; i++) {
            String line = reader.nextLine();
            arr[i] = line;
        }
        printLowestPriceForItem(arr);
    }
}
