package stringsandarrays
import"fmt"

func MarkWithZero(matrix [3][4]int,rLen int,cLen int){
   rows  := new([3]int)
   cols  := new([4]int)

   for  i := 0; i < rLen; i++ {
      for j := 0; j < cLen; j++ {
         if rows[i] != 1 && cols[j] != 1{
             fmt.Printf("Processing %d : %d\n",i,j)
             if matrix[i][j] == 0{
                setRowAndColsToZeroV2(&matrix,i,j,rLen,cLen)
                rows[i] = 1
                cols[j] = 1
             }
         } 

      }
   }

   fmt.Println("Transformed Matrix : %v",matrix)
}

func setRowAndColsToZeroV2(matrix *[3][4]int , row int ,col int,rLen int,cLen int){
  for i := 0; i < cLen;i++{
      matrix[row][i] = 0
  }  

  for i:= 0;i< rLen;i++{
     matrix[i][col] = 0
  }
}
