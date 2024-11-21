import java.io.*;
import java.util.HashMap;
import java.util.Map;
import java.util.StringTokenizer;

class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));
        StringTokenizer st;

        int indexValue = 0;
        int index = 0;

        for(int i=0 ; i<3 ; i++){
            st = new StringTokenizer(br.readLine());
            try{
                indexValue = Integer.parseInt(st.nextToken());
                index = i;
            } catch(NumberFormatException ignored){}
        }

        indexValue += 4 - (index+1);
        if(indexValue % 3 == 0 && indexValue % 5 == 0 ){
            System.out.println("FizzBuzz");
        } else if (indexValue % 3 == 0) {
            System.out.println("Fizz");
        } else if (indexValue % 5 == 0) {
            System.out.println("Buzz");
        } else {
            System.out.println(indexValue);
        }
    }
}