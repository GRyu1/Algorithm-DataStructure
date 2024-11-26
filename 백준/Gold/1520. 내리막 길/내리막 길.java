import java.io.*;
import java.util.Stack;
import java.util.StringTokenizer;

class Main {
    static int m, n;
    static int[][] route, dp;
    static int[] dx = {0, 0, 1, -1};
    static int[] dy = {1, -1, 0, 0};
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter writer = new BufferedWriter(new OutputStreamWriter(System.out));
        StringTokenizer st = new StringTokenizer(reader.readLine());
        m = Integer.parseInt(st.nextToken());
        n = Integer.parseInt(st.nextToken());
        route = new int[m][n];
        dp = new int[m][n];
        for (int i = 0; i < m; i++) {
            st = new StringTokenizer(reader.readLine());
            for (int j = 0; j < n; j++) {
                route[i][j] = Integer.parseInt(st.nextToken());
                dp[i][j] = -1;
            }
        }
        writer.write(String.valueOf(countRoute(0,0)));
        writer.flush();
        writer.close();
    }

    private static int countRoute(int x, int y) {
        if( x == m-1 && y == n-1 ) {
            return 1;
        }

        if (dp[x][y] != -1) {
            return dp[x][y];
        }

        dp[x][y] = 0;
        for (int i = 0; i < 4; i++) {
            int nx = x + dx[i];
            int ny = y + dy[i];
            if( nx < 0 || nx >= m || ny < 0 || ny >= n ) {
                continue;
            }
            if( route[nx][ny] < route[x][y] ) {
                dp[x][y] += countRoute(nx, ny);
            }
        }
        return dp[x][y];
    }
}