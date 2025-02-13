// Copyright (c) 2022 Proton AG
//
// This file is part of Proton Mail Bridge.
//
// Proton Mail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Proton Mail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Proton Mail Bridge. If not, see <https://www.gnu.org/licenses/>.

package updater

// DefaultPublicKey is the public key used to sign builds.
const DefaultPublicKey = `-----BEGIN PGP PUBLIC KEY BLOCK-----

mQINBFo9OeEBEAC+fPrLcUBY+YUc5YiMrYJQ6ogrJWMGC00h9fAv3PsrHkBz0z7c
QFDyNdNatokFDtZDX115M0vzDwk5NkcjmO7CWbf6nCZcwYqOSrBoH8wNT9uTS/6p
R3AHk1r3C/36QG3iWx6Wg4ycRkXWYToT3/yh5waE5BbLi/9TSBAdfJzTyxt4IpZG
3OTMnOwuz6eNRWVHkA48CJydWS6M8z+jIsBwFq4nOIChvLjIF42PuAT1VaiCYSmy
4sU1YxxWof5z9HY0XghRpd7aUIgzAIsXUbaEXh/3iCZDUMN5LwkyAn+r5j3SMNzk
2htF8V7qWE8ldYNVrpeEwyor0x1wMzpbb/C4Y8wXe8rP01d0ApiHVRETzsQk2esf
XuSrBCtpyLc6ET1lluiL2sVUUelAPueUQlOyYXfL2X958i0TgBCi6QRPXxbPjCPs
d1UzLPCSUNUO+/7fslZCax26d1r1kbHzJLAN1Jer6rxoEDaEiVSCUTnHgykCq5rO
C3PScGEdOaIi4H5c6YFZrLmdz409YmJEWLKIPV/u5DpI+YGmAfAevrjkMBgQBOmZ
D8Gp19LnRtmqjVh2rVdr8yc5nAjoNOZwanMwD5vCWPUVELWXubNFBv8hqZMxHZqW
GrB8x8hkdgiNmuyqsxzBmOEJHWLlvbFhvHhIedT8paU/spL/qJmWp3EB4QARAQAB
tExQcm90b24gVGVjaG5vbG9naWVzIEFHIChQcm90b25NYWlsIEJyaWRnZSBkZXZl
bG9wZXJzKSA8YnJpZGdlQHByb3Rvbm1haWwuY2g+iQJUBBMBCAA+AhsDBQsJCAcC
BhUICQoLAgQWAgMBAh4BAheAFiEE1R5k0+Y+3D7veGTO4sddaOYjSwcFAl432eEF
CQXb04AACgkQ4sddaOYjSwd9ww/9FmQa/Fh1lgE9Ug6zQMlr20UDxfCVvE+Hxn4V
OFSWLH+c491BWJMCSI/vm2XJSzjchoeYB+Ns5M/b1tC4orCzbUGb0INpcnNOZPYM
jcMlIqFlMdYzG7ZRFUX3BaMgpb0Xlyk4bLP0FcDIyJuO/53qsi4QNLNqIJOD2IDK
mG3z17GCZ+heJcttMzkzihYX6dBOeD2MUhSruTGLzGRstbVntthdpIs9u2jUCPuB
qZB2Dw2l1MtqB5UguE7Xxwz9R6xZ7a/P13wCXzVoA0Ud/pkyZ5UgAWapulBrjrCD
z4Oqa3DQpscVzex1bkj9Xd9duwBM4BbR5r5432sYiGYV1IByw8oeLQBz6APSIauN
LUxXRhKZQwqEVKigMkmofSHdQnoaEylDKKgBJRYhxpkIPY9BIup/83e8Q7wceIYM
hSJ5GvAPAMi+kWYrUgGqfUlSYNXTgswnvPWQXCHXsy9HCpFcSdsrXQr0kyZlxIGd
TSIV5hAZZL4cURXdDU+rrNJuA/Pjcebw9aTSNi/LYB+Dv4EsxIUND8d6H5bKCdeN
PFO1BXLkcwrTaOk+HNHYlwcM3H4p3MPRMCXaXXVRH1O60Rla2SGAOTuj/Xpv0Bo7
2vfNSQAEk+yHTD1iMy9IxYy/xVHbExp7ErqYARLsmw1enKCdw4h5TbL2ThBTEmje
tNYOlVyJASIEEAECAAwFAlpcl/gFAwASdQAACgkQlxC4m8pXrXzF1ggAoS7luFCm
S13Vv2w2GGpWOLcVh/RUcsTU8eUr9DY40rlrKVkX5MBL1yeD/XiIXY5aFlBaKxIq
NPjqu0VBZhaYj6ZuGpAodpattzjNOXWxwFtz2JaUfn2VUrZMbDwY9AQMHab/xxir
PmezHMee9Y56qnNPIHDh3pZZ18rHrwY4e1pVkR+N0xYTb4M0vw3AhHjboS8H9noq
V6ykT5F+3C18G5UBHwyGS/wCXf7xB7mAN4voBZq8NMe3bVae8Lk7xSCuXuzmHZPM
5q6MJB+18HSraKsFRlEJSeESb1JlKS0JnocnHxq7pdvAIw10QCC3ZF7Bu0PGDwUI
7ymZvWOsRmqiZYkBIgQQAQIADAUCWn7fTQUDABJ1AAAKCRCXELibyletfFnGB/9b
RumxnZzyoOrsDiV6DVruagouK8RhilAd+3We47l8rtSd27M8AL4RkO6JBqM7MKP/
C9anbY/2R6vRTXVF+hJ06dqek2dba3+bWi5SxmNa4Hqxp16Ip1RuH+yqdGB2gdhN
hgsY0Ojnl7vFWk2DKTzlMzP6TEXXhC4Z4XwtXRx1y7XsHen8/f5+Zo6ro565KuD+
RuE+6WQQ0h6yhtEsuMeohNSLYpqZo9d2hBE09768gdweSSB+4FyIQsIBWjLX/iyB
WiBxw3LtkQwl21TGbntYD12Gr0hJRY7c3meg/PN+XKYTcAml7BOvvdaEpWLVfs/h
W9QOicuL0l/74GZ1GFKuiQEiBBABAgAMBQJakANDBQMAEnUAAAoJEJcQuJvKV618
SwgH/jF/S4jfpKbwid8aigJs8CSSL3GQFtjU5/6qiMUJQD1BC9WpVMZImm+8y+qK
MzTWR479o6GRChq4YPCkzvK74/lbGLacugsBtVkRzvDRcHVNUjl9RhUdxvU1Wwsr
ASSIZdLOMXWpaQhxNgrkM2DDLX+mWWExwsHbuOS0DfFeQeVmtgmfJgWb2bhc9X0V
ZBfcDmQ8F2Fazkf25E+PozfTKxMbCcj4Hzht8eWRGqsvJbM+Guf/7P2GXq69clD2
h7TDdCyTvHWyruCKqrHjYRUXbE8U5j8FnEZh6nvH+1OfF6Pt7SNhMDjeGzcI2ddl
MDBO/EGNPEpyL3Eeh0n7xuE1vL6JASIEEAECAAwFAlqzm54FAwASdQAACgkQlxC4
m8pXrXxCgwf/VsUKgIn5BANRu6tHCuk68aT7gj7RiO3F1Ta2170sy3/hguXP9k+k
dO1wHaILkN5h6ge/Ant+mSO8Vod6nfEakBSfaPfdXf1Wa5fTu0rUI9L9PV2lgTs3
N6R7C0YQ3lDyylvX85cfZbel3n0aSr1XFb1FFPl7CeWy97Qnx3XMHbLI6uiALNK9
8GcYUA/lXWzDfGv86O6n5/d9K2q7QA0XW95IDegy7Jacchtp8AjHuZ6xa8ADFYFL
qNIoK+/4PH0p/piiNHrQ3Ndys31Kpi2X/TrVPhZ0OLtUk1qUdtSLK3fwPXstuSrp
04dyVWRWa/a22Qap+4/JLmGP4J6DUmUisYkBIgQQAQIADAUCWsVn8gUDABJ1AAAK
CRCXELibyletfAdTB/9VNysmCsCD1tV54h45iU91GWy9GzoiqKQ2aKPzHX9CG4uo
GwWSKUj0cMwoqvw0ysAJ78S1G11N3DR4j98PzlcJ7s+jXUB57PC3Va8dhajyjbp7
hbNE2jrgqYQyTp/XcHd2xJWqQtniRtY1bruFP/0HbflAq3t4Y63xTjtM6kj2xi0t
wauvOSzPAvIb7zJj8lmLmzOZ+cCuOfZJG37QrLefMztLQAq2676VQr3wjBU4tcbk
FJctn5cb7VIR/act/aW0mutnPF3fBKO8d0ILFj2j72cuWL5dFlWu6biF/GR2hodE
gQpIcDAfaOKxM2XK7Ii8wBizqgZAo6vVBGsJgC6TiQEiBBABAgAMBQJbC0a8BQMA
EnUAAAoJEJcQuJvKV618gjIH/ib3CEeXjKb5unUoZTSRUiHNRyL8WBnF6jTD8zw3
+8SkBWTZQnlO/29HTU5hth99yG5VoN2wooVusYZuPMXUEoR1DHpPRzR4JeZ2TmON
sB6siXYqpoO4TkSCh0utCzm3SADSiie8rq0ijWNiuooVBfFTiyrrJ4TifS1jP58t
CWWkmb0JcO41jVtGeGLDeYfTcR7iiuYh5EddHeqw4d1WwwE9VnYyy7inR/yyBCFF
I+zHyQ7IXVLlzJkKOIZEWdYsSsbA4LXTaVNRHP73UlONPJLVEdbcgm4GG20WLrFj
eH0E7RQABypCFFZMAovqKhrcc2DuGQjSb1TF3trRp43L9ceJAbMEEAEKAB0WIQQc
kUEqiCSCZfLZBb/QmabVW538jAUCW2q2/wAKCRDQmabVW538jDemC/9ZpRxijeVX
8LjouNaVOh0+TJfQbEpZOIBuoP88m2O9jZoEsRiRLMyd4+3v8TiYZYobPrVZ5/Cl
pX33XmblRq3y2FKnnI8cKKiKqGKpmscV2IbGR53GaV6DYfqTva/sCmAQmKeyLuvo
+t5I7SN1I32vathkvlMxq0YH79PSG3BYYASLOEg9D8eyqKn8DBdsw9uuKnXdzBFB
T+UjqO3Pw8+pD6D2bSQSSYldKTCxwtiZFl2TtrCyWVM03rP3lKSOIx+9xNNw1T4X
mpbflYejFpWj+pAjmTI3Qfy4O5e582F8gUnrlZ3g3R/7jMOmKXw7xpQikFgtfFYo
MZjBd9bs8LhCdyi9KnVLeq3Svd+HaVLb7KS2pk9bcXvpZBja7A1F2U5yBX7dGGg5
1kCgNb/FXhmjb2MrHCNSPCqks4nzvUzsdviI4Q+gjYgZMaytj0uRmAe3bsC1WnMt
QFI+hjKbay92m9OUCen1nwtwxKrv1JeSN08id3UlRK2Y2yyt3NwZtOOJAlQEEwEI
AD4CGwMFCwkIBwIGFQgJCgsCBBYCAwECHgECF4AWIQTVHmTT5j7cPu94ZM7ix11o
5iNLBwUCW/fvvAUJA7ze2wAKCRDix11o5iNLB+GeD/9Y2FmqvRK243gpth1Ab4Vj
5ouKMuGJUjeLiPrBSKz7tRYU0xQZ4+wFSsXvM6vOjhWnXoVJGGGm94hsMxUmDjIa
XPA36Nmv0UG8XzORs6qNqORMXes7KpJAKllWB5qZG3pziAZBM7B/DEHgrmZWDiu/
rkhIgqMtS0JnbuCkPp4mzbkuHJlCcjrkkekcORVqvuhDadta/S2fbCutQoET2FEz
kmfXEZodiStjE94c0MeurElQNs98mFxrngz31uz1C6bcqRPVZZTn/S2sDRXx5Rlr
3fPJ0s3KcklypuopRgG86enMv/SXeF19+aiWlRdYWuUU+U4tVCUrgpzqhCpLYpYH
JJAeD2SVO2jWn1yk1kTew7n5R+V7JdbR03oNxT/GLChgon/7Hglli2+of+Q/uC97
ajOcSwszDqTZBmg7UaKESmhek57Ozjr4RH9gbHhkXad7ZuZifOg45kFsZfCwxlM1
pr1AkdtzVvP2OOMgUFzXOat2LSexuOW8u6ARghtJv9Y/TArAbrrNNq2yVEhDa6eV
Qk2dZvQThUuDP3KrntCm/FLsPfXZ1lXQ2cDCHIMj7nhcCK9dNBvR8AZ9yBu0p0qk
lvV1wAxv1Y4GemsErtcuxZsmPm/mcHcC5Z6vs2FTIkH+iBZC2pFRjVsBwjhwHXMX
Bir88vSxt8AcPHFkQXpeGIkCVAQTAQgAPhYhBNUeZNPmPtw+73hkzuLHXWjmI0sH
BQJaPTnhAhsDBQkB4TOABQsJCAcCBhUICQoLAgQWAgMBAh4BAheAAAoJEOLHXWjm
I0sHLZIQAIovDkggSMkgjxUn92ZNwTTR8KwKM2tKy9EItpWJl5p0j/5mXFfNsDg7
R93sJnimrS1bOSAiJHN5P+I0gNXOME8pP55B+oM3ttHJbOfUb30gVktlvNILhFxZ
W+TO//LD5KK20TupRe31GOVN6hF9h0WY+EhX2W3tFMVTy48BakwxRa2EbBHLRhE3
Uvo+I+tFSnRNpMyZSo1+Kj4ZGI7tKBNzW9QNTSCB06fhrC3SeAfn3lHCTmbJG42Z
DNY5Yyc9XdBzFoPwXu+kQJ01aI46SwpPjc1tn6K4TJm6mDhyGuOlQfBBoPrr/MOf
UFLJVxviZwy0XKQSaYepzvZDYPAnpfVQ5ig4XR+SN1bYpQua31TZTDuCmRGEUtTU
ZC8xCvzzLWsKa4U6VnPeMwcg3B6vL4r24wFb3hIQej1T7xaLkrTFzWBH8v7vLmQc
qRSvpkU05N9LP+M/C6Q4CMzgPdv3QFHQ5fKiBp/Csk8ZqJb+UsCKhRB/HBDkhxLK
NlMIQTXG7fdeSei6BE39EWqz/BpiJqPPsl9nrnA1nv4EGZefeq6U/fPJcMwszV4S
atGpU+STaU7OvIELSUV/N6KlU9QuFlFM69GfVncfC1K0bOXyA0hN3nakurOKEa6K
Uno8kcDyJkgUnnvmPKh/PmSEa3fRxK/InuCJftl1q3OfqHa0QEv9iQEiBBABAgAM
BQJchPx5BQMAEnUAAAoJEJcQuJvKV618VpkH/j51WqzA0b6SWMFu7vzTV2kSQduO
fgLpYmHSsAQPdeipnIbYQWftYxJ9obbRWjieVyO58g2aiJnorKPNcM0qh2XKszUk
dK5h0930SmDz86yFP4MVeAHIsZVub5c/oXC665IDzp4cqDIGftaX8xomIt9eoUOO
c+Zzp3/yW2wXd8tvpSMhnAU3RDK269DsjF+rz8OAcU5KufxtwyDzxNrha+xHrhp+
NLvCuWdbCvhrpeRWj46EkRfiSNZ2hAUWi6KgPdCR8bVD9eYfSIXa1HrSHHJQOb/Z
mPkF5voi21OPmeH88+WmnkYjUrKgD3b1toETsWgrT6iKmeomsmd+Dx5esCCJASIE
EAECAAwFAlyWyLcFAwASdQAACgkQlxC4m8pXrXwNiwf9Hd87+e2Nmg3QrXGhcTJg
lR0BW8x5FpKekYEJ6rYcbNsyd5mz0SU0FplU924HNgu9m4E8wSpvN6gUEbjMmp0M
yJBohxiyI+Ii5ZqZdg29GYX5IHxFn3x8IfVeOTXR4rtABwidl7eOeUT085iy81Vt
iiB0O/DHf3NhPInCtWkcXuQwEZGnm+Gjb3cpURtg7iOdU7gZFqzvfRt8dPLOJ3+c
nRikseh3bGscB8d3FloX9Yu4a/7QyAXgzcz28GWOhfbQ305Rtr6pnyXEEpaJL5De
xjYPtEVjbRvbFvTGHHlWRovkQLpdXXjXNY57efjIRVbASqpb19LgdwdyMXlcBrlN
xIkBIgQQAQgADAUCXKiU9wUDABJ1AAAKCRCXELibyletfLZsB/4mRQOQ3qnXOzvz
82ZBo1j5XxYzlwHL5qeVqxyyVvbq4obQmO6T4lABD1Fdn6WfqjnP+gCsnapCFC2e
UswxYUyt2m6EWREAsPHaacCsRqgL4FAZKIgdhlFkv8op0gUhs2++n892Asse9Nx7
ZAvkXJiC52LQjxO/HyD07+JmjHjQvvNYr9Lwrli1jqzNQaAYB7zgkxPUGVmLYdsQ
iaMNZq24NTahGwVzxZThZkdN34gOuazxWAxAqYkEmEvM7TucB8HQovxlUsUgw5yI
KkhMsZriZE69yhPMwby7mSJa7k3xjGchTkojKcszLRi+e1HgtTDVbD7fudGaHa0i
NDhEEygxiQEiBBABCAAMBQJcumE2BQMAEnUAAAoJEJcQuJvKV618Oi0IAKUTTb1Y
HIjPe1EAi5lEhnemClZTxhEDsoMJplV8X/hTpsByVhSZa+P7Lgdplko8r7iAXXJl
87hNiZensPB4SMM+/j5pQi+H8rbo03sZUzKmvcn0JT0cFM7sn/g7Q4aDOI9k78kd
XZA7lFRubn2j7QVQVS3BXMTIf8/vFqredK5z0diyLMC/6sCi+l1vAk2Kbf7e4v/X
7HETaDCfqHpIQOv0/VaS3q1NpmSeaYM9fCLOHi0KByEt7oGG7vX/SXkUCgVRF6T6
4oIHHBskm5nJVUzeZTPV9x+bVBy6svJHRkMMHI5bsrlGTnCDJHBTQMMRmef6D+Pj
qMfN6m38q5FGvw25Ag0EWj054QEQALdPQOlRT1omHljxnN64jFuDXXSIb6zqaBvU
wdYoDpV2dfRmzGklsCVA7WHXBmDWbUe9avgO3OO7ANw6/JzzYjP+jwImpJg7cSqT
qW8A1U6TYfGXVUV3a/obIEttl7bI9BsUNgmLsBYIwHov+gl/ajKQdALYHCmq3Bj6
o7BBeWPpVpk9dzjcsLVbmNszNGP1Ik5dKE0jZUi6h+YoVuJE9o/+T+jxoqFRpXNs
ZqWOEKmCHDz6TTs1iTp+CoZ/5g0eKph6XJ+TuNoqF9491IYEFn9oxzsoIBkewTY/
fJWmXf++cnpBODrZLF/GoRFc7MW9Kael9vmQ0J7mjM2bFs308lH0rRrfmdlLAU5i
KgPv0akxnnnUqvCcoekFMURDtP3z09KZXuOMnt834utd7WLe+LZD6dxs+rPhyDiW
80E8Bdlz1Jo+c2g6toIN+uD7/f5gwaZaXhJB0oO7fWSVVo+HJprWBnmf9frgKq1O
cS0BNvA+4Aip2hhFqWJAbUQXCyMaeU2WTWIzy0FQ6SEFFy/RM8O5O1HHsDYjtIic
9QJ/PqSD0qN7LMlkjR8AdWvAxm95i5GpxDZODldsOneeummvsn3I1jCoULTik7iJ
VdRuY1V3vfsYAkefGN/n2ga3MvatCJipwoCGsMgUXGTdokXOqKBgMBuBLCkxj2wl
ol2R9p8RABEBAAGJAjwEGAEIACYCGwwWIQTVHmTT5j7cPu94ZM7ix11o5iNLBwUC
XjfaBAUJBdvTowAKCRDix11o5iNLByTCD/oCRk97JjCqNb4B1Ed/G5tJ+w55cptq
1dBZOxvEf+ol/403Q+R5bRqun3vXYupzZyIEvi10OVZ/t3t/FboOAWwJ222o0Ivm
t6RhErlmF1dCsKILy5i0iLJexLFGJIiSh6Slr2BZoiqasrlCYStJE2hXhNjXOIiZ
76YsLed6b5MKBllsw4DGPgT9sECrWft935oGo9caVUTX5VsnoVvzxKQLYki8m1Et
Eki7M3MK2pPNpX6y1e862JvL0oUfjYjrn8ALrgTeNtx/oRDgMujD1UQd5kGdwzkG
ec1nB7T5Gdiyfd8unDvSd+Eg3UgF3eDgFA8ZDdO7yZlWv3aEeVUUAvEDT9/RgbnZ
a22GhGcCJ3mHBMbx2khLIorJq8C1ZkhzpKIWqmETgr3MvUo+iT9YsnuGd8qpl2JK
Ru2QuKZ0VTqLMuURMojMETiRbfxBg8uZMAxPr45Hqq2hj/8ooF/hYS/Y2oD+b8DM
7hSTTEXm14p3tp8BbFGdVq1jJn5Zsj7isLAzydlyWWKcfwcljpzLCOo/wME3zUVh
mDPZYW3/sexJl/ROUHXo8eqBEMUgNLjffiuymfh7L8RmVOcsJsS87nu+iVvR5CaJ
0VLGn+SuxFT09xhvM4NxQIgNgk+NuQeIcwOMd6vtvf0btSTflN3hRGhGzLcZWWww
m//Hk8dcT8vncg==
=G/D6
-----END PGP PUBLIC KEY BLOCK-----`
