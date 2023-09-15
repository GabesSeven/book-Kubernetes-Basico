#  Cria uma fila de tarefas chamada 'keygen'
curl -X PUT localhost:8080/memq/server/queues/keygen

# Cria 100 itens de tarefas e os carrega na fila
for i in work-item-{0..99}; do
  curl -X PUT localhost:8080/memq/server/queues/keygen/enqueue -d "$i"
done
