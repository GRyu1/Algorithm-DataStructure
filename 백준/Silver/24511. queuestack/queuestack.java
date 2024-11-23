import java.io.*;
import java.util.*;

class Main {
    static int n,m;
    static int[] arr;
    static Deque<Integer> data = new ArrayDeque<>();
    static int[] inputs;
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new OutputStreamWriter(System.out));
        StringTokenizer stringTokenizer;
        n = Integer.parseInt(new StringTokenizer(bufferedReader.readLine()).nextToken());
        arr = new int[n];

        stringTokenizer = new StringTokenizer(bufferedReader.readLine());
        int i = 0;
        while(stringTokenizer.hasMoreTokens()){
            arr[i++] = Integer.parseInt(stringTokenizer.nextToken());
        }

        stringTokenizer = new StringTokenizer(bufferedReader.readLine());
        i = 0;
        while(stringTokenizer.hasMoreTokens()){
            int a = Integer.parseInt(stringTokenizer.nextToken());
            if(arr[i++]==0) data.addLast(a);
        }

        m = Integer.parseInt(new StringTokenizer(bufferedReader.readLine()).nextToken());
        inputs = new int[m];
        stringTokenizer = new StringTokenizer(bufferedReader.readLine());
        i = 0;
        while(stringTokenizer.hasMoreTokens()){
            inputs[i++] = Integer.parseInt(stringTokenizer.nextToken());
        }

        for ( i=0 ; i < m ; i++) {
            data.addFirst(inputs[i]);
            bufferedWriter.write(data.pollLast()+" ");
        }
        bufferedWriter.flush();
        bufferedWriter.close();
    }
}

