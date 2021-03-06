
Έ
google/protobuf/empty.protogoogle.protobuf"
EmptyBv
com.google.protobufB
EmptyProtoPZ'github.com/golang/protobuf/ptypes/emptyψ’GPBͺGoogle.Protobuf.WellKnownTypesJώ
 3
Μ
 2Α Protocol Buffers - Google's data interchange format
 Copyright 2008 Google Inc.  All rights reserved.
 https://developers.google.com/protocol-buffers/

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are
 met:

     * Redistributions of source code must retain the above copyright
 notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above
 copyright notice, this list of conditions and the following disclaimer
 in the documentation and/or other materials provided with the
 distribution.
     * Neither the name of Google Inc. nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


  

" ;
	
%" ;

# >
	
# >

$ ,
	
$ ,

% +
	
% +

& "
	

& "

' !
	
$' !

( 
	
( 
ϋ
 3 ο A generic empty message that you can re-use to avoid defining duplicated
 empty messages in your APIs. A typical example is to use it as the request
 or the response type of an API method. For instance:

     service Foo {
       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
     }

 The JSON representation for `Empty` is empty JSON object `{}`.



 3bproto3
Ύ=
environment.proto!endpoints.terrariumai.environmentgoogle/protobuf/empty.proto"
Entity
id (	RidI
classID (2/.endpoints.terrariumai.environment.Entity.ClassRclassID
x (Rx
y (Ry
energy (Renergy
health (Rhealth
ownerUID (	RownerUID
modelID (	RmodelID"1
Class	
EMPTY 	
AGENT
ROCK
FOOD"ω
Effect
x (Rx
y (Ry
	timestamp (R	timestampI
classID (2/.endpoints.terrariumai.environment.Effect.ClassRclassID
value (Rvalue
decay (Rdecay
	delThresh (R	delThresh" 
Class
NONE 
	PHEROMONE"X
CreateEntityRequestA
entity (2).endpoints.terrariumai.environment.EntityRentity"&
CreateEntityResponse
id (	Rid""
GetEntityRequest
id (	Rid"V
GetEntityResponseA
entity (2).endpoints.terrariumai.environment.EntityRentity"%
DeleteEntityRequest
id (	Rid"0
DeleteEntityResponse
deleted (Rdeleted"Υ
ExecuteAgentActionRequest
id (	Rid[
action (2C.endpoints.terrariumai.environment.ExecuteAgentActionRequest.ActionRactiond
	direction (2F.endpoints.terrariumai.environment.ExecuteAgentActionRequest.DirectionR	direction"1
Action
WAIT 
MOVE
EAT

ATTACK"2
	Direction
UP 
DOWN
LEFT	
RIGHT"Ύ
ExecuteAgentActionResponsea
value (2K.endpoints.terrariumai.environment.ExecuteAgentActionResponse.ResponseValueRvalue"=
ResponseValue
OK 
ERR_INVALID_TARGET
ERR_DIED"8
GetEntitiesInRegionRequest
x (Rx
y (Ry"d
GetEntitiesInRegionResponseE
entities (2).endpoints.terrariumai.environment.EntityRentities"7
GetEffectsInRegionRequest
x (Rx
y (Ry"a
GetEffectsInRegionResponseC
effects (2).endpoints.terrariumai.environment.EffectReffects2Σ
Environment
CreateEntity6.endpoints.terrariumai.environment.CreateEntityRequest7.endpoints.terrariumai.environment.CreateEntityResponse" x
	GetEntity3.endpoints.terrariumai.environment.GetEntityRequest4.endpoints.terrariumai.environment.GetEntityResponse" 
DeleteEntity6.endpoints.terrariumai.environment.DeleteEntityRequest7.endpoints.terrariumai.environment.DeleteEntityResponse" 
ExecuteAgentAction<.endpoints.terrariumai.environment.ExecuteAgentActionRequest=.endpoints.terrariumai.environment.ExecuteAgentActionResponse" >

ResetWorld.google.protobuf.Empty.google.protobuf.Empty" =
	SpawnFood.google.protobuf.Empty.google.protobuf.Empty" 
GetEntitiesInRegion=.endpoints.terrariumai.environment.GetEntitiesInRegionRequest>.endpoints.terrariumai.environment.GetEntitiesInRegionResponse" 
GetEffectsInRegion<.endpoints.terrariumai.environment.GetEffectsInRegionRequest=.endpoints.terrariumai.environment.GetEffectsInRegionResponse" JΗ'
  

  

 *
	
  %
 
   Taks we have to do



 
5
  ( Unique integer identifier of the agent


  

  

  	

  

  
 Entity stats


  


   

   	

   

  

  	

  

  

  

  

  

  

  

 

 

 

 

 

 

 

 

 	


 

 

 

 

 	


 

 

 

 

 	

 

 

 

 

 	

 

  Owner details


 

 

 	

 

 

 

 

 	

 
U
 -I Effects are manipulations of a cell that last a specific amount of time





  Pos


 

 

 	


 







	




  Time of creation


 

 

 

 

 "% Class


 "

  #

  #

  #

 $

 $

 $

&

&%

&

&

&
*
( Value, used for observation


(&

(

(	

(
'
* How fast it should decay


*(

*

*

*
'
, Min before being deleted


,*

,

,	

,
.
0 3" Request data to create new agent



0

 2 agent


 20

 2

 2	

 2
,
5 8  Contains data of created agent



5
"
 7 ID of created agent


 75

 7

 7	

 7
)
; > Request data to read entity



;
5
 =( Unique integer identifier of the agent


 =;

 =

 =	

 =
=
@ C1 Contains entity data specified in by ID request



@
%
 B Task entity read by ID


 B@

 B

 B	

 B
*
F I Request data to delete agent



F
?
 H2 Unique integer identifier of the agent to delete


 HF

 H

 H	

 H
1
K O% Contains status of delete operation



K
c
 NV Contains number of entities have beed deleted
 Equals 1 in case of successful delete


 NK

 N

 N

 N


Q d


Q!
>
 S1 Id for the agent that should perform the action


 SQ#

 S

 S	

 S
)
 UZ identifier for the action


 U

  V

  V

  V

 W

 W

 W

 X

 X

 X


 Y

 Y


 Y

[

[Z

[

[	

[
2
]b$ direction to perform the action in


]

 ^

 ^

 ^	


_

_

_

`

`

`

a

a	

a

c

cb

c

c

c


	e m


	e"

	 gk
 Response


	 g

	  h

	  h

	  h	


	 i

	 i

	 i

	 j

	 j

	 j

	 l

	 lk

	 l

	 l

	 l



o r



o"


 p


 po$


 p


 p	



 p


q


qp


q


q	



q


s u


s#

 t

 t


 t

 t

 t


w z


w!

 x

 xw#

 x

 x	


 x

y

yx

y

y	


y


{ }


{"

 |

 |


 |

 |

 |
,
   Service to manage simulation


 
 
  I Create new agent


  

  &

  1E
&
 @ Get data for an entity


 

  

 +<

 I Delete an agent


 

 &

 1E
0
 -  Perform an action for an agent


 

 2

 )

 J Reset the world


 

 &

 1F
'
 I Spawn food in the world


 

 %

 0E

 ^ Get Region


 

 4

 ?Z

 [

 

 2

 =Wbproto3