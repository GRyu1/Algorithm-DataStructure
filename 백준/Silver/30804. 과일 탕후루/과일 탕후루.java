import java.io.*;
import java.util.*;

class Main {
    static Map<Integer, Integer> map;
    static int[] arr;
    static int n;
    static int count = 0;
    static int max = 0;
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));
        n = Integer.parseInt(br.readLine());
        arr = new int[n];
        map = new HashMap<>();
        StringTokenizer st = new StringTokenizer(br.readLine());
        for (int i = 0; i < n; i++) {
            arr[i] = Integer.parseInt(st.nextToken());
        }
        
        int left = 0;
        for(int right = 0 ; right < n ; right++){
            int temp = map.getOrDefault(arr[right], 0);
            map.put(arr[right], temp + 1);
            if (temp == 0) {
                count++;
            }

            while(count > 2){
                temp = map.get(arr[left]);
                if(temp == 1){
                    map.remove(arr[left]);
                    count --;
                } else {
                    map.put(arr[left], temp - 1);
                }
                left ++;
            }
            max = Math.max(max, right - left + 1);
        }
        bw.write(max + "");
        bw.flush();
        bw.close();
    }
}

