CGO_ENABLED=1
CGO_LDFLAGS='-static -s'
CGO_CXXFLAGS='-static-libgcc -static-libstdc++ -Wl,-Bstatic -lstdc++ -lpthread -Wl,-Bdynamic'

run1:
	go run ./cmd/1_uso_camara.go

clean:
	rm -rf $(TARGET)

.PHONY: run1 clean

.DEFAULT: 
	@echo 'No hay disponible ninguna regla para este destino'
