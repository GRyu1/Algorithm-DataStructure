import java.io.*;
import java.util.ArrayList;
import java.util.List;
import java.util.StringTokenizer;

class Main {
    static int m, n;
    static int[] arr1, arr2;
    static List<boolean[]> dp;

    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder sb = new StringBuilder();

        StringTokenizer st = new StringTokenizer(reader.readLine());
        n = Integer.parseInt(st.nextToken());
        arr1 = new int[n];

        st = new StringTokenizer(reader.readLine());
        int i = 0;
        while(st.hasMoreTokens()) {
            arr1[i++] = Integer.parseInt(st.nextToken());
        }

        st = new StringTokenizer(reader.readLine());
        m = Integer.parseInt(st.nextToken());
        arr2 = new int[m];

        st = new StringTokenizer(reader.readLine());
        i = 0;
        while(st.hasMoreTokens()) {
            arr2[i++] = Integer.parseInt(st.nextToken());
        }

        dp = new ArrayList<>();
        for(i = 0; i <= n; i++) {
            dp.add(new boolean[500*(i+1)+1]);
        }

        solve(0,0);

        for (i = 0; i < m; i++) {
            if (arr2[i] > 500 * 30) {
                sb.append("N ");
            } else if (dp.get(n)[arr2[i]]) {
                sb.append("Y ");
            } else {
                sb.append("N ");
            }
        }
        System.out.println(sb);
    }

    private static void solve(int i, int target) {
        if (i > n) {
            return;
        }

        if (dp.get(i)[target]) {
            return;
        }

        dp.get(i)[target] = true;
        solve(i + 1, target);
        if (i < n) {
            solve(i + 1, target + arr1[i]);
            solve(i + 1, Math.abs(target - arr1[i]));
        }
    }
}