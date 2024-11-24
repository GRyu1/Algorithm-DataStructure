import java.io.*;
import java.util.*;

class Main {
    static List<Integer> lis;
    static int[] arr;
    static int n;
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        n = Integer.parseInt(br.readLine());
        arr = new int[n];
        lis = new ArrayList<>();
        StringTokenizer st = new StringTokenizer(br.readLine());
        for (int i = 0; i < n; i++) {
            arr[i] = Integer.parseInt(st.nextToken());
        }
        lis.add(arr[0]);

        for (int i = 1; i < n; i++) {
            if (lis.get(lis.size() - 1) < arr[i]) {
                lis.add(arr[i]);
            } else {
                int idx = binarySearch(arr[i]);
                lis.set(idx, arr[i]);
            }
        }
        System.out.println(lis.size());
    }

    static int binarySearch(int target) {
        int left = 0;
        int right = lis.size() - 1;
        while (left <= right) {
            int mid = (left + right) / 2;
            if (lis.get(mid) == target) {
                return mid;
            } else if( lis.get(mid) < target ){
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return left;
    }
}