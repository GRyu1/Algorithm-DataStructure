import java.io.*;
import java.util.*;

class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new OutputStreamWriter(System.out));
        StringTokenizer stringTokenizer;
        stringTokenizer = new StringTokenizer(bufferedReader.readLine());
        int capacity = Integer.parseInt(stringTokenizer.nextToken());

        Deque<int[]> deque = new ArrayDeque<>(capacity);
        stringTokenizer = new StringTokenizer(bufferedReader.readLine());
        int k=0;
        while(stringTokenizer.hasMoreTokens()){
            int datum=Integer.parseInt(stringTokenizer.nextToken());
            deque.addLast(new int[]{datum, ++k});
        }

        for (int i = 0; i < capacity-1; i++) {
            int[] displacement = deque.poll();
            bufferedWriter.write(displacement[1]+" ");
            if( displacement[0] > 0 ){
                for(int j=0; j < displacement[0]-1 ; j++){
                    deque.addLast(deque.pollFirst());
                }
            } else {
                int distance = Math.abs(displacement[0]);
                for(int j=0; j < distance ; j++ ){
                    deque.addFirst(deque.pollLast());
                }
            }
        }
        bufferedWriter.write(deque.poll()[1]+"");
        bufferedWriter.flush();
        bufferedWriter.close();
    }
}