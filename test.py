import hashlib

class HashableFileStruct:
    def __init__(self, fileName, fileSize, mimeType, fileID):
        self.fileName = fileName
        self.fileSize = fileSize
        self.mimeType = mimeType
        self.fileID = fileID

    def pack(self):
        hasher = hashlib.md5()
        for field in [self.fileName, str(self.fileSize), self.mimeType, str(self.fileID)]:
            hasher.update(field.encode())
        return hasher.hexdigest()

def pack_file(fileName, fileSize, mimeType, fileID):
    return HashableFileStruct(fileName, fileSize, mimeType, fileID).pack()

print(pack_file(
		"",
		12863576,
		"video/mp4",
		6066524942051052953,
)[:6])