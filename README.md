# Go Design Patterns

### Injeção de dependência
Injeção de dependência é um padrão de projeto que nos permite injetar uma dependência em um objeto. Essa dependência não precisa ter uma implementação, podendo ser apenas um contrato/interface.

Esse padrão nos permite ter uma maior flexibilidade em relação a essa dependência, podendo alterar a implementação sem a necessidade de alterar o código que a utiliza.

Geralmente queremos injeção de dependência quando se trata de comunicação com serviços externos, como banco de dados, APIs e etc. Dessa forma temos maior flexibilidade para alterar o provedor de serviço sem a necessidade de alterar o código que o utiliza. Esse mesmo padrão nos permite realizar testes unitários, por exemplo, já que podemos injetar um mock da dependência.


### Proxy 
Proxy é um padrão de projeto que nos permite criar uma representação de um objeto. Dessa forma, podemos criar uma representação de um objeto e controlar o acesso a ele.  
É como adicionar um objeto intermediário entre o cliente e o objeto real. 

Geralmente usamos proxy quando queremos controlar o acesso a um determinado objeto, podendo adicionar lógica antes e depois da execução do objeto real. 


### Facade 
Facade é um padrão de projeto que nos permite criar uma interface simplificada para um conjunto de interfaces complexas. Dessa forma, podemos ter uma interface mais simples para utilizar, sem a necessidade de conhecer as interfaces complexas por trás.

Geralmente usamos facade quando queremos simplificar a utilização de um determinado serviço, possibilitando ao cliente executar uma única ação (única função) para realizar uma determinada tarefa.







