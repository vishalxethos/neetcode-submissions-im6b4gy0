func trap(height []int) int {

        // maxLeft := 0
        // maxRight := 0

        // result := 0

        // leftMax := []int{}
        // rightMax := make([]int,len(height))

        // for _,elem := range height{

        //     leftMax = append(leftMax,maxLeft)
        //     maxLeft = max(maxLeft,elem)



        // }

        // for i := len(height) - 1; i >= 0 ; i -- {
        //     rightMax[i] = maxRight
        //     maxRight = max(maxRight, height[i]) 
        // }


        // for index := range leftMax {
        //     result += max(min(leftMax[index],rightMax[index]) - height[index],0)
        // }

        // return result

        maxLeft := 0
        maxRight := 0
        result := 0

        L,R := 0, len(height) - 1

        for L < R {
            
            if height[L] < height[R] {
                result += max(maxLeft - height[L],0)
                maxLeft = max(maxLeft,height[L])
                L+=1
            } else {
                result += max(maxRight - height[R],0)
                maxRight = max(maxRight, height[R])
                R-=1

            }


        }

        return result
}
