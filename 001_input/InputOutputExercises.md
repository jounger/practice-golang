Bài tập xử lý input/output console và đọc file:

Yêu cầu:

1. Cho phép input từ console
2. Input này là đường dẫn đến 1 file config.json
3. Nếu không phải file json thì hiển thị invalid file, chọn lại đường dẫn
4. Nếu đúng là file json xử lý tiếp, chưa cần validate xem json có đúng định dạng hay không
5. Giả sử ban đầu không có instances nào được bật.
6. Sau khi đọc file config, in ra consoles theo mẫu:
	[<tên của instance type>] tab [<hành động>] tab [<số lượng>] \n
	[<tên của instance type>] tab [<hành động>] tab [<số lượng>] \n
	...
	
ví dụ: 
	["a1.medium"] [provision] [2]
	["a2.large"]  [delete]	  [6]
	
trong đó: tên của instance types tham khảo tại, tên của instance type sẽ được lưu ở json 
	https://aws.amazon.com/ec2/instance-types/
	
	hành động: 2 hành động chính: provision và delete. Provision: nếu cần thêm máy thì bật máy
	delete: nếu cần tắt bớt máy thì tắt bớt.
	
	Ví dụ như hiện đang có 10 instances medium. file config nhận được là giờ chỉ cần 5 instances medium.
	hiển thị ra sẽ là: tắt đi 5 instances
	
7. Sau khi in ra console, tiếp tục cho phép nhập đường dẫn đến file config khác	
	