package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
//    "container/list"
)

var count int=0;

func
sort( str[] rune, start , end rune )[] rune{

            var i rune =0; 
            var j rune= 0;

            for i=start; i <= end ; i++ {

                    for j=start ; j <= end ;j++{

                            if str[i] <= str[j]{

                                str[i],str[j] = str[j],str[i]

                            }

                    }

            }

           return str

}



func index( str[] rune , first , l , h rune) rune {

        var index rune = l;
        var i rune = 0;
        var value rune =first;

        for i = l+1 ; i<= h ; i++ {

                        if str[i] > value && str[i] < str[index] {

                                        index = i;
                        }

        }

        return index

}



func permutation( str []rune ) { 


        var size rune = rune(len(str));

        //fmt.Println("size = ", size );

        var fin rune = 0;
        count = 0;

        str = sort(str, 0, size-1)

        //fmt.Println( "str =", string(str))

        for fin == 0 {

              count++;

             // fmt.Println("str = " ,string(str) );

             // fmt.Println("count :", count)

              var i rune = 0;

              for i = size-2; i >= 0 ; i-- {

                  if str[i] < str[i+1] {

                               break;

                   }

              }


               if i == -1 {

                    fin = 1;

               }else{

                     var index rune = index( str , str[i] , i+1 , size-1 );

                     str[i],str[index] = str[index] , str[i]

                     var length rune = rune(len(str))

                     length = length - 1;

                     str = sort(str, i+1, length ); //rune(len(str))-i-1)
              }

    }


}



func return_permutaion( str string ) int {

     var str_array [] rune;

     str_array = []rune(str[1:len(str)]);


     fmt.Println("str = ", string(str_array));

     permutation(str_array );


     return count;

}



func handle_event( w http.ResponseWriter, r *http.Request ){

fmt.Println(" request path = %s " , r.URL.Path );
fmt.Println(" escape path = %s " , html.EscapeString(r.URL.Path) );
fmt.Fprintf(w, "Hi, Permiutaion of str =  %d", return_permutaion(r.URL.Path))


}

func main() {


    http.HandleFunc("/", handle_event );

    log.Fatal(http.ListenAndServe(":8080", nil))

}
