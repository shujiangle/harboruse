package harborget

import (
	"fmt"
	"harboruse/pkg/harborget"
	"log"
	"os"
	"strings"
)

func Harborgetmain(url string, username string, password string, harborfile string) {

	// 去掉url最后/
	url = strings.TrimSuffix(url, "/")

	client := harborget.NewClient(url, username, password)

	projects, err := client.GetProjects()
	if err != nil {
		log.Fatalf("Error fetching projects: %v", err)
	}

	// 使用 os.OpenFile 打开文件，如果文件不存在则创建，如果文件已存在则返回错误
	// 打开文件以附加内容
	file, err := os.OpenFile(harborfile, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	for _, project := range projects {
		fmt.Println("=============================================================================")
		projectdesc := fmt.Sprintf("Project ID: %d, Name: %s\n, Create: %v\n", project.ProjectID, project.Name, project.CreationTime)
		fmt.Println(projectdesc)
		repositories, err := client.GetRepositories(project.ProjectID)
		if err != nil {
			log.Fatalf("Error fetching repositories for project ID  %v", err)
		}

		_, _ = file.WriteString("=============================================================================\n")
		_, _ = file.WriteString(projectdesc)
		_, _ = file.WriteString("\n")

		for _, repo := range repositories {

			repositoriestag, err := client.GetRepositoriesTag(repo.Name)
			if err != nil {
				log.Fatal("%s存在", repositoriestag)
			}

			for _, tag := range repositoriestag {
				tagdesc := fmt.Sprintf("拉取地址: docker pull %v/%s:%s\n", url, repo.Name, tag.Name)
				fmt.Println(tagdesc)
				_, _ = file.WriteString(tagdesc)
			}

		}

		defer file.Close()
	}
}

func Harborgetprojectmain(url string, username string, password string, harborfile string) {
	// 去掉url最后/
	url = strings.TrimSuffix(url, "/")

	client := harborget.NewClient(url, username, password)

	projects, err := client.GetProjects()
	if err != nil {
		log.Fatalf("Error fetching projects: %v", err)
	}

	// 使用 os.OpenFile 打开文件，如果文件不存在则创建，如果文件已存在则返回错误
	// 打开文件以附加内容
	//file, err := os.OpenFile(harborfile, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
	//if err != nil {
	//	fmt.Println("Error opening file:", err)
	//	return
	//}

	for _, project := range projects {
		fmt.Printf("%s\n", project.Name)
	}
}
