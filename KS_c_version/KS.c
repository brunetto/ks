/*nome del file:posizioni_finali_sup2.c
data:6/11/2002
questo programma calcola la posizione finale proiettata del CM della BS
nell'ipotesi in cui la linea di vista coincida con l'asse z*/
#include <stdio.h>
#include <stdlib.h>
#include <math.h>

#include "nrutil.h"
#include "nr.h"

#define SQUARE(X) ((X)*(X))

FILE *vel1, *vel2, *vel3;

main(){
  
  int i;
  int N1=398; //number of objects in distribution 1
  int N2=520; //number of objects in distribution 2
  double distr1[N1], distr2[N2];
  double d, prob;


  vel1=fopen("BSS_scl_ardece.dat","r"); //file containing objects in distribution 1
  vel2=fopen("BSS_scl_ardecw.dat","r");//file containing objects in distribution 2
  vel3=fopen("KS_prob.out","w"); //output KS statistics
 
  for(i=0; i<N1; i++){
    fscanf(vel1, "%lf", &distr1[i]);
  }

  for(i=0; i<N2; i++){
    fscanf(vel2, "%lf", &distr2[i]);
  }

  kstwo(distr1,N1,distr2,N2,&d,&prob);

  fprintf(vel3, "%12.7e %12.7e\n", d, prob); //d = D factor in KS statistics (see Fig. 14.3.1 p.624 NRecipes in C, prob = probability of finding a D larger than observed --> if prob = 1 the two distr. are the same, if prob = 0 have nothing to do

   return 0;
}
