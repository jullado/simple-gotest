## วิธีการสั่ง run test
- สั่งรันโดยใช้คำสั่ง go test โดยต้องเข้าไปที่ dir ที่มี package _test นั้นก่อน
- สั่งรันโดยใช้คำสั่ง go test <module_name>/<package_name> (ไม่ต้องใส่ _test)
- สั่งรันทุก packages โดยใช้คำสั่ง go test ./...

### options run test
- ใส่ -v เพื่อดู function ที่ถูก test
- ใส่ -cover เพื่อดูว่าเราเทสครอบคลุมไปแล้วกี่ % ของทุกเคสที่อาจเกิดใน function ที่เราเทส (ตั้งค่าเช็ค auto ใน setting user JSON ได้)
- ใส่ -run="<func_test/sub_test>" เพื่อเลือกเทสเฉพาะ sub test นั้นๆ

- ใส่ -bench เพื่อให้รัน func benchmark ด้วย โดยใส่ -bench=. เพื่อรันทั้งหมด และ -bench=<func_bench> เพื่อเลือกรัน
- ใส่ -benchmem ต่อจาก -bench เพื่อแสดงการใช้ memory
- รูปแบบ benchmark แสดงเป็น logical-core, จน.รอบ (operation), time/รอบ, Byte/รอบ, allocate mem/รอบ


## ตั้งค่า setting user JSON เพื่อแสดง covered แบบ auto

(...เพิ่มต่อจากการตั้งค่าเดิม...)
```json
"go.coverOnSave": true, // เช็คว่าเราเขียนเทสครอบคลุมทุกเคสยังตอน save (ขึ้นไฮไลท์เขียวแดง)
"go.coverageDecorator": {   
    "type": "gutter",   // กรณีไม่ชอบแบบไฮไลท์ ให้ไปอยู่แทบข้างๆแทน
    "coveredHighlightColor": "rgba(64,128,128,0.5)",
    "uncoveredHighlightColor": "rgba(128,64,64,0.25)",
    "coveredBorderColor": "rgba(64,128,128,0.5)",
    "uncoveredBorderColor": "rgba(128,64,64,0.25)",
    "coveredGutterStyle": "blockgreen", // สี covered
    "uncoveredGutterStyle": "blockred"  // สี uncovered
},
```
