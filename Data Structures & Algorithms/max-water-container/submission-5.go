func maxArea(heights []int) int {
        L,R := 0, len(heights) - 1

        maxArea := 0

        for L < R {
                currentArea := min(heights[L],heights[R]) * (R - L)

                maxArea = max(maxArea, currentArea)

                if heights[L] <= heights[R] {
                    L +=1
                } else {

                    R-=1

                }
        }

        return maxArea


}
