package constant

var PERMISSION = []map[string]interface{}{
	{
		"TYPE":           "BIND_ACCESSIBILITY_SERVICE",
		"SECURITY_LEVEL": "Signature",
		"DESCRIPTION":    "สำหรับใช้งานกับผู้พิการ ซึ่งส่งผลให้รับข้อมูลการแตะหน้าจอ สร้าง Overlay เหนือหน้าจอแอปพลิเคชัน และอื่น ๆ รวมทั้งสิ้น 22 รายการ",
	},
	{
		"TYPE":           "REQUEST_DELETE_PACKAGES",
		"SECURITY_LEVEL": "Normal",
		"DESCRIPTION":    "ลบแอปพลิเคชันได้",
	},
	{
		"TYPE":           "QUERY_ALL_PACKAGES",
		"SECURITY_LEVEL": "Normal",
		"DESCRIPTION":    "ค้นหาแอปพลิเคชันที่อยู่บนเครื่องได้",
	},
	{
		"TYPE":           "GET_INSTALLED_APPS",
		"SECURITY_LEVEL": "Normal",
		"DESCRIPTION":    "ดูรายชื่อของแอปพลิเคชันอื่นที่ติดตั้งในเครื่อง",
	},
	{
		"TYPE":           "ACCESS_NETWORK_STATE",
		"SECURITY_LEVEL": "Normal",
		"DESCRIPTION":    "เข้าถึงสถานะของอินเตอร์เน็ต",
	},
	{
		"TYPE":           "FOREGROUND_SERVICE",
		"SECURITY_LEVEL": "Normal",
		"DESCRIPTION":    "เมื่อแอปพลิเคชันเริ่มทำงาน สามารถย้ายการทำงานไปเป็นรูปแบบ Foreground ได้ โดยเมื่ออยู่ในรูปแบบนี้เมื่อผู้ใช้งานปิดหน้าจอของระบบ ระบบจะยังคงทำงานอยู่ เช่น การเปิดแอปพลิเคชันเล่นเพลง แล้วปิดหน้าแอปพลิเคชันไป เพลงยังคงเล่นต่อ ",
	},
	{
		"TYPE":           "READ_PHONE_STATE",
		"SECURITY_LEVEL": "Dangerous",
		"DESCRIPTION":    "เข้าถึงข้อมูล เช่น ข้อมูล Cellular สถานะของการโทรศัพท์",
	},
	{
		"TYPE":           "WRITE_EXTERNAL_STORAGE",
		"SECURITY_LEVEL": "Dangerous",
		"DESCRIPTION":    "สร้างไฟล์ในเครื่องตามตำแหน่งที่แอปพลิเคชันติดตั้งได้",
	},
	{
		"TYPE":           "CAMERA",
		"SECURITY_LEVEL": "Dangerous",
		"DESCRIPTION":    "แอปพลิเคชันเข้าถึงกล้องได้",
	},
	{
		"TYPE":           "RECORD_AUDIO",
		"SECURITY_LEVEL": "Dangerous",
		"DESCRIPTION":    "แอปพลิเคชันอัดเสียงได้",
	},
	{
		"TYPE":           "INTERNET",
		"SECURITY_LEVEL": "Normal",
		"DESCRIPTION":    "อนุญาตให้ใช้งานอินเตอร์เน็ตได้",
	},
	{
		"TYPE":           "ACCESS_WIFI_STATE",
		"SECURITY_LEVEL": "Normal",
		"DESCRIPTION":    "ดึงข้อมูลเกี่ยวกับ Wifi เช่น ชื่อ ความแรงสัญญาณ ",
	},
	{
		"TYPE":           "CHANGE_NETWORK_STATE",
		"SECURITY_LEVEL": "Normal",
		"DESCRIPTION":    "แก้ไขการเชื่อมต่อ ทั้ง Wifi และ Mobile data",
	},
	{
		"TYPE":           "READ_SMS",
		"SECURITY_LEVEL": "Dangerous",
		"DESCRIPTION":    "เข้าถึงข้อมูลข้อความ เช่น รายละเอียดข้อความ ผู้ส่ง ผู้รับ หมายเลขโทรศัพท์ เวลา",
	},
	{
		"TYPE":           "RECEIVE_SMS",
		"SECURITY_LEVEL": "Dangerous",
		"DESCRIPTION":    "เข้าถึงข้อมูลข้อความขณะที่ข้อความเข้าได้ทันที",
	},
	{
		"TYPE":           "SEND_SMS",
		"SECURITY_LEVEL": "Dangerous",
		"DESCRIPTION":    "ส่งข้อความได้ รวมทั้งดูได้ว่าข้อความถูกส่งหรือไม่",
	},
	{
		"TYPE":           "SYSTEM_ALERT_WINDOW",
		"SECURITY_LEVEL": "Signature",
		"DESCRIPTION":    "อนุญาตให้แอปพลิเคชันแสดงกล่องข้อความได้ไม่ว่าจะปิดอยู่หรือไม่ ",
	},
	{
		"TYPE":           "SYSTEM_OVERLAY_WINDOW",
		"SECURITY_LEVEL": "Dangerous",
		"DESCRIPTION":    "แสดงผลในลักษณะหน้าต่างเหนือแอปพลิเคชันอื่นได้ รวมทั้งหน้าต่างของระบบ",
	},
	{
		"TYPE":           "MOUNT_UNMOUNT_FILESYSTEMS",
		"SECURITY_LEVEL": "Dangerous",
		"DESCRIPTION":    "เข้าถึงไฟล์ทั้งในเครื่องและจากภายนอก ",
	},
	{
		"TYPE":           "READ_EXTERNAL_STORAGE",
		"SECURITY_LEVEL": "Dangerous",
		"DESCRIPTION":    "อ่านไฟล์ที่อยู่ใน SD cards ได้",
	},
	{
		"TYPE":           "MODIFY_AUDIO_SETTINGS",
		"SECURITY_LEVEL": "Dangerous",
		"DESCRIPTION":    "ปรับการตั้งค่าของเสียงของระบบได้ เช่น เพิ่ม ลด เสียง เปิด ปิด เสียง รวมทั้งเลือก Output ของเสียงได้",
	},
	{
		"TYPE":           "WRITE_SETTINGS",
		"SECURITY_LEVEL": "Dangerous",
		"DESCRIPTION":    "ปรับการแก้ไขการตั้งค่าของระบบได้ เช่น ความสว่างหน้าจอ ระดับเสียง เสียงรอสาย",
	},
	{
		"TYPE":           "ACCESS_NOTIFICATION_POLICY",
		"SECURITY_LEVEL": "Normal",
		"DESCRIPTION":    "เปิด ปิด การแจ้งเตือนได้",
	},
	{
		"TYPE":           "ACCESS_FINE_LOCATION",
		"SECURITY_LEVEL": "Dangerous",
		"DESCRIPTION":    "ดึงข้อมูลตำแหน่งในระดับแม่นยำ",
	},
	{
		"TYPE":           "WAKE_LOCK",
		"SECURITY_LEVEL": "Normal",
		"DESCRIPTION":    "ให้หน่วยประมวลผลทำงานได้จากสถานะพักหน้าจอ",
	},
	{
		"TYPE":           "DISABLE_KEYGUARD",
		"SECURITY_LEVEL": "Normal",
		"DESCRIPTION":    "สามารถเข้าถึงทรัพยากรโดยไม่จำเป็นต้องยืนยันรหัสผ่านของเครื่อง",
	},
	{
		"TYPE":           "REQUEST_INSTALL_PACKAGES",
		"SECURITY_LEVEL": "Signature",
		"DESCRIPTION":    "สามารถเปิดตัวติดตั้งเพื่อติดตั้งแอปพลิเคชันได้",
	},
	{
		"TYPE":           "REQUEST_IGNORE_BATTERY_OPTIMIZATIONS",
		"SECURITY_LEVEL": "Normal",
		"DESCRIPTION":    "ข้ามการจำกัดการใช้งานแบตเตอรี่",
	},
	{
		"TYPE":           "xxx",
		"SECURITY_LEVEL": "xxx",
		"DESCRIPTION":    "xxx",
	},
	{
		"TYPE":           "xxx",
		"SECURITY_LEVEL": "xxx",
		"DESCRIPTION":    "xxx",
	},
}
