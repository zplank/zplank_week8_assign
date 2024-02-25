# zplank_week8_assign

The statistical method chosen for this assignment was counterfactual causal inference, which involves utilizing statistics to understand the causal relationship between variables and understanding what could happen under different circumstances. A few examples of packages in R that can be used for this are:
    - Randomized Controlled Trials (RTCs) using 'randomizeR' or 'randomizr'
        - https://www.rdocumentation.org/packages/randomizeR/versions/3.0.2
        - https://www.rdocumentation.org/packages/randomizr/versions/1.0.0
    - Propensity Score Matching using 'MatchIt'
        - https://www.rdocumentation.org/packages/MatchIt/versions/4.5.5
    - Regression Models using 'lm' or 'glm' 
        - https://www.rdocumentation.org/packages/stats/versions/3.6.2/topics/lm
        - https://www.rdocumentation.org/packages/stats/versions/3.6.2/topics/glm
    - Causal Inferences of causal effects in regression modeling using 'CausalImpact'
        - https://www.rdocumentation.org/packages/CausalImpact/versions/1.3.0

------------------------

For the purposes of this assignment, I will be testing functionality of creating Linear Regression Models in R (using lm) and Go using a dataset I have on hand. This input data incudes variables associated with homes that sold and aims to predict SalePrice based on explanatory variables such as square foot, room count, lot area, and more. See the 'regression_model_R.r' file to see how this was completed in R and comparatively the regression_model_go.go file to see how this was completed in Go. 

------------------------

Overall, when comparing R and Go to for this specific instance of building counterfactual causal inference, both programs produced similar results and had similar processing statistics when producing a linear regression model (see appendix of summary results below). That being said, each had their pros and cons throughout the process. My initial reaction to comparing the two programs is R works in a more simple fashion when it comes to importing data and building a regression model. It requires very little code (one line for each) in order to produce what you are looking for, however, Go allows for this process to be more repeatable and once established and set up properly, the code can be used across various input datasets. R was also easier to use when it came to automated variable selection. I was not able to even produce a working model using a stepwise regression method in Go (could be lack of familiarity with this specific process but I tried multiple options and was failing to produce code) but in R, it is a relatively easy process, again only requiring a line or two of code and very little manual manipulation needed for omitted missing values or handling different data types. 

For this specific project, my recommendation would be to continue using R. I find it much more user friendly, produces solid results, and has quick processing times. That being said, Go would likely be a better choice for a project that is more involved and data clean up, pre-processing, and variable selection is not as big of a challenge. 


-----------------------
Appendix - Summary Results 

R 
Residuals:
    Min      1Q  Median      3Q     Max 
-557683  -18013   -2205   16279  275979 

Coefficients:
               Estimate Std. Error t value Pr(>|t|)    
(Intercept)  -7.716e+04  6.740e+03 -11.448  < 2e-16 ***
TotalBsmtSF   1.937e+01  3.471e+00   5.580 2.70e-08 ***
FirstFlrSF    5.821e+01  4.705e+00  12.372  < 2e-16 ***
SecondFlrSF   3.825e+01  3.877e+00   9.866  < 2e-16 ***
LotArea       7.583e-01  1.355e-01   5.597 2.45e-08 ***
OverallQual   2.456e+04  8.790e+02  27.944  < 2e-16 ***
BsmtFullBath  1.768e+04  1.709e+03  10.349  < 2e-16 ***
BsmtHalfBath  1.804e+03  3.384e+03   0.533   0.5939    
FullBath      1.345e+04  2.082e+03   6.459 1.29e-10 ***
HalfBath      9.714e+03  2.109e+03   4.605 4.35e-06 ***
BedroomAbvGr -7.733e+03  1.448e+03  -5.340 1.03e-07 ***
KitchenAbvGr -3.735e+04  4.708e+03  -7.934 3.30e-15 ***
TotRmsAbvGrd  4.199e+03  1.033e+03   4.063 5.00e-05 ***
Fireplaces    3.722e+03  1.478e+03   2.519   0.0118 *  
---
Signif. codes:  0 ‘***’ 0.001 ‘**’ 0.01 ‘*’ 0.05 ‘.’ 0.1 ‘ ’ 1

Residual standard error: 37780 on 2255 degrees of freedom
Multiple R-squared:  0.796,	Adjusted R-squared:  0.7948 
F-statistic: 676.7 on 13 and 2255 DF,  p-value: < 2.2e-16

> cat("Processing Time:", processing_time[3], "seconds\n")
Processing Time: 0.02 seconds



Go
Regression Formula: Predicted = -77155.2807 + X0*19.3659 + X1*58.2092 + X2*38.2542 + X3*0.7583 + X4*24563.2607 + X5*17682.6281 + X6*1804.4139 + X7*13447.4992 + X8*9713.8610 + X9*-7732.9452 + X10*-37351.6753 + X11*4198.5338 + X12*3722.3221
R^2: 0.795968
Coefficients: 0x76c3a0
Processing Time: 1.1690942s



