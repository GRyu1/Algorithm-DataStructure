import java.io.*;
import java.util.*;

class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        List<Integer> data = new ArrayList<>();
        Deque<Integer> deque = new ArrayDeque<>();

        int count;
        count = Integer.parseInt(new StringTokenizer(bufferedReader.readLine()).nextToken());
        StringTokenizer st;
        for(int i=0; i<count; i++){
            st = new StringTokenizer(bufferedReader.readLine());
            while(st.hasMoreTokens()){
                data.add(Integer.parseInt(st.nextToken()));
            }
        }

        for(int i=0 ; i< data.size() ; i++){
            switch(data.get(i)){
                case 1:
                    deque.addFirst(data.get(++i));
                    break;
                case 2:
                    deque.add(data.get(++i));
                    break;
                case 3:
                    if(deque.size()>0){
                        System.out.println(deque.pollFirst());
                    } else {
                        System.out.println(-1);
                    }
                    break;
                case 4:
                    if(deque.size()>0){
                        System.out.println(deque.pollLast());
                    } else {
                        System.out.println(-1);
                    }
                    break;
                case 5:
                    System.out.println(deque.size());
                    break;
                case 6:
                    System.out.println(deque.size() == 0 ? 1 : 0);
                    break;
                case 7:
                    if(deque.size()>0){
                        System.out.println(deque.peekFirst());
                    } else {
                        System.out.println(-1);
                    }
                    break;
                case 8:
                    if(deque.size()>0){
                        System.out.println(deque.peekLast());
                    } else {
                        System.out.println(-1);
                    }
                    break;
            }
        }
    }
}