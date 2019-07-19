/**
 * 
 */
package com.komal.taskexecutor.task_example;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.io.Reader;
import java.net.URL;
import java.nio.charset.Charset;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

import org.apache.commons.io.FileUtils;
import org.tautua.markdownpapers.ast.CodeSpan;
import org.tautua.markdownpapers.ast.Document;
import org.tautua.markdownpapers.ast.Header;
import org.tautua.markdownpapers.ast.Node;
import org.tautua.markdownpapers.ast.Paragraph;
import org.tautua.markdownpapers.parser.ParseException;
import org.tautua.markdownpapers.parser.Parser;

/**
 * @author kkadam
 *
 */
public class MarkDownParser {
	public static void downloadWithApacheCommons(String url, String localFilename) {

        int CONNECT_TIMEOUT = 10000;
        int READ_TIMEOUT = 10000;
        try {
            FileUtils.copyURLToFile(new URL(url), new File(localFilename), CONNECT_TIMEOUT, READ_TIMEOUT);
        } catch (IOException e) {
            e.printStackTrace();
        }

    }


	/**
	 * @param args
	 */
	public static void main(String[] args) {
		List<String> attributeList = new ArrayList<>();
		//String resource_name = "ec2_capacity_reservation";
		//String resource_name = "ec2_client_vpn_endpoint" + ".html";
		//String markdown_file_name = "security_center_subscription_pricing";
		String markdown_file_name = args[0];
		String provider_type = args[1];
		String released_version = args[2];

		parseAndGetAttributeList(attributeList, provider_type, released_version, markdown_file_name);
		writeAtributesToFile(attributeList);
		
	}


	private static void parseAndGetAttributeList(List<String> attributeList, String provider_type, String released_version,
			String markdown_file_name) {
		String url = getMarkdownURL(provider_type, released_version, markdown_file_name);
		String localFilename = markdown_file_name + ".markdown";
		try {
			downloadWithApacheCommons(url, localFilename);
		} catch (Exception e) {
			/*url = getMarkdownURL(provider_type, released_version, resource_name + ".html");
			downloadWithApacheCommons(url, localFilename);*/
		}
		
		Reader in;
		try {
			in = new FileReader(localFilename);
			Parser parser = new Parser(in);

			Document doc = parser.parse();
			int argumentNodeNumber = 0;
			for (int i=0; i < doc.jjtGetNumChildren(); i ++) {
				Node docChild = doc.jjtGetChild(i);
				if (docChild instanceof Header) {
					String header = "";
					Node headerChild = docChild.jjtGetChild(0);
					for (int j=0; j < headerChild.jjtGetNumChildren(); j ++) {
						header = header +  headerChild.jjtGetChild(j);
					}
					if (header.trim().equals("Argument Reference"))
						argumentNodeNumber = i;
				}
				
			}
			Node AttrubuteChilds = doc.jjtGetChild(argumentNodeNumber + 2);
			if (AttrubuteChilds != null) {
				int totalchildren  = AttrubuteChilds.jjtGetNumChildren();
				for (int i = 0; i < totalchildren; i++) {
					Node ithChild = AttrubuteChilds.jjtGetChild(i);
					for (int j = 0; j < ithChild.jjtGetNumChildren(); j++) {
						//String string = ithChild[j];
						Node itemChild = ithChild.jjtGetChild(j);
						if (itemChild instanceof Paragraph) {
							String attributeName = "";
							String attributeDesc = "";
							Node firstChild = itemChild.jjtGetChild(0);
							for (int k = 0; k < firstChild.jjtGetNumChildren(); k++) {
								Node paragraphChild = firstChild.jjtGetChild(k);
								if (paragraphChild instanceof CodeSpan && k==0) {
									attributeName = ((CodeSpan) paragraphChild).getText();
									attributeList.add(attributeName);
								} else {
									if (paragraphChild instanceof CodeSpan)
										attributeDesc = attributeDesc +  ((CodeSpan) paragraphChild).getText();
									else
										attributeDesc = attributeDesc + paragraphChild.toString();
								}
							}
							System.out.println(attributeName + " : " +attributeDesc);
						} else {
							System.out.println("Item child : " +itemChild.toString());
						}
						//System.out.println("Item child : " +itemChild.toString());
						
					}					
					//((Item)ithChild).
					//System.out.println(ithChild);
					
				}
			}
		} catch (FileNotFoundException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		} catch (ParseException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		new File(localFilename).delete();
	}


	private static void writeAtributesToFile(List<String> attributeList) {
		String attribute_list_name = "../attributes.go";
		String attribute_array = "package main\nvar Attributes_array = []string{";
		for (String attribute : attributeList) {
			attribute_array = attribute_array + "\"" +attribute + "\"";
			if (attributeList.indexOf(attribute) != attributeList.size() - 1)
				attribute_array = attribute_array + ", ";
		}
		attribute_array = attribute_array + "}\n";
		try {
			FileUtils.writeStringToFile(new File(attribute_list_name), attribute_array, Charset.defaultCharset());
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
	}


	private static String getMarkdownURL(String provider_type, String released_version, String resource_name) {
		return "https://raw.githubusercontent.com/terraform-providers/terraform-provider-" + provider_type + "/" 
				+ released_version + "/website/docs/r/" + resource_name + ".markdown";
	}

}
