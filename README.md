# redlineBodyFile

While testing [Fireeye's Redline](https://fireeye.market/apps/211364), I noticed that the 'files-api.urn-*UUID*.xml' file has enough information to allow creating a body file so you can create a timeline.  This program will allow specifiying a path from the <FullPath> tag for each FileItem (see the example below).

`````
<?xml version="1.0" encoding="UTF-8"?>
<itemList generator="files-api" generatorVersion="30.19.0" itemSchemaLocation="http://schemas.mandiant.com/2013/11/fileitem.xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
 <FileItem created="2023-09-11T02:36:57Z" uid="UUID">
  <FullPath>/lib64</FullPath>
  <FilePath />
  <FileName>lib64</FileName>
  <FileExtension />
  <SizeInBytes>9</SizeInBytes>
  <Modified>2022-02-06T15:05:00Z</Modified>
  <Accessed>2023-09-11T00:34:45Z</Accessed>
  <Changed>2022-02-06T15:05:00Z</Changed>
  <Username>root</Username>
  <SecurityID>0</SecurityID>
  <Group>root</Group>
  <GroupID>0</GroupID>
  <Permissions>777</Permissions>
  <FileAttributes>Symlink</FileAttributes>
 </FileItem>
</itemList>
`````
## Running the program:

Help option:

    redlineBodyFile -h
    -----
    Usage of redlineBodyFile:

	-d string
	      The directory to scan (no trailing slash) or full file path.
	-f string
	      The RedLine Audit file.
       
Example:

    redlineBodyFile -f redline_audit.xml -d "c:\\documents"

Depending on the size of the XML file, this may take a while to run.

Then you can parse it with your normal tool to create a timeline.

    
