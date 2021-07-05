Decorator pattern
===

Decorator는 progarming pattern의 하나이다.

볼펜이 있다고 할 때 볼펜의 기본기능인 쓰기기능만 가지고 있는 볼펜도 있지만 볼펜 케이스에 캐릭터를 그려 꾸민다던지 볼펜에 라이트를 단다던지 이와 같이 기본기능 이외에 이것 저것 꾸미는 것과 같이 프로그램에서도 기본적인 기능에 부가적인 기능들을 추가하는 design pattern이라고 볼 수 있다.

즉 기본 기능 이외에 추가되는 기능들을 Decorator라고 한다.

그럼 왜 사용이 되는 걸까?

부가기능들은 대체적으로 잘 바뀐다. 이러한 부가기능과 기본기능을 하나의 소스파일에 만들게 되면 부가기능이 많아질 수록 코드 유지보수에 힘들어지고 SOLID원칙에 어긋나게 된다.(Open Close principle) 그렇기 때문에 부가기능들을 Decorator로 구현을 함으로 기본기능은 Code수정 없이 변경되는 부가기능만 변경해주면 되기 때문이다.

- Decorator

<img src = "https://user-images.githubusercontent.com/74294325/124489923-7a5f0500-ddec-11eb-9751-2fd665da9c3f.png">

다이어그램을 보면 Component를 구현한 두 개의 구조체가 있다. Component1은 기본기능을 구현한 Component이고 Decorator는 Component interface를 Member변수로 가지고 있다.

Client는 Component interface를 가지고 있다. (구체적으로는 Decorator1의 인스턴스를 가지고 있다. 가장 앞에 있는 Decorator 인스턴스)

호출 과정은 Client는 Decorator1을 호출한다. 그 후 Decorator2의 operator를 호출한다. Decorator2는 Component1의 operator를 호출한다. 이 후 역순으로 작업이 진행되고 마지막으로 Decorator1의 operator의 작업까지 완료 후 return받게 된다.