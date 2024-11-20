import java.io.*;
import java.util.StringTokenizer;

class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new OutputStreamWriter(System.out));
        
        int sum = 0;
        int n = Integer.parseInt(bufferedReader.readLine().trim());

        StringTokenizer stringTokenizer = new StringTokenizer(bufferedReader.readLine());
        int[] sizeAmount = new int[stringTokenizer.countTokens()];
        for (int i = 0; i < sizeAmount.length; i++) {
            sizeAmount[i] = Integer.parseInt(stringTokenizer.nextToken());
        }

        stringTokenizer = new StringTokenizer(bufferedReader.readLine());
        int[] p = new int[stringTokenizer.countTokens()];
        for (int i = 0; i < p.length; i++) {
            p[i] = Integer.parseInt(stringTokenizer.nextToken());
        }

        for (int value : sizeAmount) {
             sum += value % p[0] == 0 ? value/p[0] : value/p[0]+1;
        }
        bufferedWriter.write(sum+"\n");
        bufferedWriter.write(n/p[1] + " " + n%p[1]);
        bufferedWriter.flush();
        bufferedWriter.close();
    }
}