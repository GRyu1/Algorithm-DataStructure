import java.io.*;
import java.util.*;

class Main {
    static int n,m;
    static boolean[] isVisited;
    static List<List<Integer>> graph;
    static int[] dp;
    static int[] answer = new int[2];
    static int max = -1;
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));
        StringTokenizer st = new StringTokenizer(br.readLine());
        n = Integer.parseInt(st.nextToken());
        m = Integer.parseInt(st.nextToken());

        dp = new int[n+1];
        graph = new ArrayList<>();
        for(int i=0; i<=n; i++) {
            graph.add(new ArrayList<>());
        }
        for(int i=0; i<m; i++) {
            st = new StringTokenizer(br.readLine());
            int a = Integer.parseInt(st.nextToken());
            int b = Integer.parseInt(st.nextToken());
            // b를 해킹 하면 a도 해킹 됨.
            graph.get(b).add(a);
        }


        for (int i = 1; i <= n; i++) {
            isVisited = new boolean[n+1];
            solve(i);
        }

        for(int i=1; i<=n; i++) {
            if(dp[i] == max) {
                bw.write(i + " ");
            }
        }
        bw.flush();
        bw.close();
    }

    public static void solve(int start) {
        Queue<Integer> q = new LinkedList<>();
        q.add(start);
        isVisited[start] = true;

        while(!q.isEmpty()) {
            int current = q.poll();
            for(int next : graph.get(current)) {
                if(isVisited[next]) continue;
                q.add(next);
                dp[start]++;
                isVisited[next] = true;
            }
        }
        if(dp[start]==0) dp[start]++;
        if(dp[start]>max) max = dp[start];
    }
}