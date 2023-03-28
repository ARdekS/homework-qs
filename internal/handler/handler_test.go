package handler

// func TestRenameItem(t *testing.T) {
// 	// Создаем mock-объекты для параметров функции
// 	requestBody := strings.NewReader(`{"id": 2, "text": "text"}`)

// 	request := httptest.NewRequest("POST", "http://localhost:81/path", requestBody)
// 	responseRecorder := httptest.NewRecorder()

// 	// Вызываем функцию обработки POST-запроса
// 	(responseRecorder, request)

// 	// Проверяем результаты
// 	if responseRecorder.Code != http.StatusOK {
// 		t.Errorf("Expected status code %d, but got %d", http.StatusOK, responseRecorder.Code)
// 	}

// 	expectedResponse := `{"result": "ok"}`
// 	if responseRecorder.Body.String() != expectedResponse {
// 		t.Errorf("Expected response body %s, but got %s", expectedResponse, responseRecorder.Body.String())
// 	}
// }
